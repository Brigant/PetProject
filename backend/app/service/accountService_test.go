package service

import (
	"errors"
	"testing"
	"time"

	"github.com/Brigant/PetPorject/backend/app/core"
	"github.com/golang-jwt/jwt"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_CreateUser(t *testing.T) {
	type mockBehavior func(s *MockAccountStorage, account core.Account)

	testCasesTable := map[string]struct {
		account              core.Account
		mockBehavior         mockBehavior
		expectedResult       string
		expectedErrorMessage string
		wantError            bool
	}{
		"Succes": {
			account: core.Account{
				Phone:    "+3809999999",
				Password: "qwerty123456",
				Age:      15,
				Role:     "admin",
			},
			mockBehavior: func(s *MockAccountStorage, account core.Account) {
				s.EXPECT().InsertAccount(account).Return("id-111", nil)
			},
			expectedResult:       "id-111",
			expectedErrorMessage: "",
			wantError:            false,
		},
		"Error occures": {
			account: core.Account{
				Phone:    "+3809999999",
				Password: "qwerty123456",
				Age:      15,
				Role:     "admin",
			},
			mockBehavior: func(s *MockAccountStorage, account core.Account) {
				s.EXPECT().InsertAccount(account).Return("", errors.New("XXX"))
			},
			expectedResult:       "",
			expectedErrorMessage: "service CreateUser get an error: XXX",
			wantError:            true,
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			// Init Deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			AccountStorage := NewMockAccountStorage(ctrl)
			testCase.mockBehavior(AccountStorage, testCase.account)

			accountService := AccountService{
				storage: AccountStorage,
			}

			result, err := accountService.CreateUser(testCase.account)

			assert.Equal(t, testCase.expectedResult, result)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedErrorMessage)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_GenerateAccessToken(t *testing.T) {
	testCasesTable := map[string]struct {
		account core.Account
		session core.Session
	}{
		"success generate AccessToken": {
			account: core.Account{
				ID:   "ID-111",
				Role: "admin",
			},
			session: core.Session{
				RefreshToken: "RefreshToken-111",
				AccountID:    "ID-111",
				Role:         "admin",
				RequestHost:  "example.com",
				UserAgent:    "Some mozilla agent",
				ClientIP:     "127.0.0.1",
				Expired:      time.Now().Add(5 * time.Second),
				Created:      time.Now(),
			},
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			var accountService AccountService

			accessToken, _ := accountService.generateAccessToken(testCase.session)

			token, _ := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errInvalidSigningMethod
				}

				return []byte(signingKey), nil
			})

			claims, ok := token.Claims.(*Claims)
			if !ok {
				t.Error(errWrongTokenClaimType.Error())
			}

			assert.GreaterOrEqual(t, claims.ExpiresAt, claims.IssuedAt,
				"ExpiresAt cat't be less or equal to IssueAt")
			assert.Equal(t, testCase.account.ID, claims.Info.AccountID)
			assert.Equal(t, testCase.account.Role, claims.Info.Role)
			assert.Equal(t, testCase.session.RefreshToken, claims.Info.RefreshToken)
			assert.Equal(t, testCase.session.RequestHost, claims.Info.RequestHost)
			assert.Equal(t, testCase.session.UserAgent, claims.Info.UserAgent)
			assert.Equal(t, testCase.session.ClientIP, claims.Info.ClientIP)
		})
	}
}

func TestService_RefreshTokenpair(t *testing.T) {
	type mockBehaviorInsert func(s *MockAccountStorage, session core.Session)
	type mockBehaviorRefresh func(s *MockAccountStorage, session core.Session)

	testCasesTable := map[string]struct {
		session              core.Session
		behaviorInsert       mockBehaviorInsert
		behaviorRefresh      mockBehaviorRefresh
		expectedRefreshToken string
		expectedErrorMessage string
	}{
		"Refresh tokens successfull": {
			session: core.Session{
				AccountID:    "Some-ID",
				Role:         "admin",
				RefreshToken: "RefreshToken-111",
				RequestHost:  "example.com",
				UserAgent:    "Some mozilla agent",
				ClientIP:     "127.0.0.1",
			},
			behaviorInsert: func(s *MockAccountStorage, session core.Session) {
				s.EXPECT().SelectSession(session).Return(core.Session{
					RefreshToken: session.RefreshToken,
					AccountID:    session.AccountID,
					Role:         session.Role,
					RequestHost:  session.RequestHost,
					UserAgent:    session.UserAgent,
					ClientIP:     session.ClientIP,
					Expired:      time.Now().Add(5 * time.Minute),
					Created:      time.Now().Add(-5 * time.Minute),
				}, nil)
			},
			behaviorRefresh: func(s *MockAccountStorage, session core.Session) {
				s.EXPECT().RefreshSession(gomock.Any()).Return(nil)
			},
			expectedRefreshToken: "RefreshToken-111",
		},
		"RefreshToken expired": {
			session: core.Session{
				AccountID:    "Some-ID",
				Role:         "admin",
				RefreshToken: "RefreshToken-111",
				RequestHost:  "example.com",
				UserAgent:    "Some mozilla agent",
				ClientIP:     "127.0.0.1",
			},
			behaviorInsert: func(s *MockAccountStorage, session core.Session) {
				s.EXPECT().SelectSession(session).Return(core.Session{
					RefreshToken: session.RefreshToken,
					AccountID:    session.AccountID,
					Role:         session.Role,
					RequestHost:  session.RequestHost,
					UserAgent:    session.UserAgent,
					ClientIP:     session.ClientIP,
					Expired:      time.Now().Add(-5 * time.Minute),
					Created:      time.Now(),
				}, nil)
			},
			behaviorRefresh: func(s *MockAccountStorage, session core.Session) {
			},
			expectedRefreshToken: "",
			expectedErrorMessage: "refresh token has expired",
		},
		"No session in DB": {
			session: core.Session{
				AccountID:    "Some-ID",
				Role:         "admin",
				RefreshToken: "RefreshToken-111",
				RequestHost:  "example.com",
				UserAgent:    "Some mozilla agent",
				ClientIP:     "127.0.0.1",
			},
			behaviorInsert: func(s *MockAccountStorage, session core.Session) {
				s.EXPECT().SelectSession(session).Return(core.Session{}, errors.New("no session"))
			},
			behaviorRefresh: func(s *MockAccountStorage, session core.Session) {
			},
			expectedRefreshToken: "",
			expectedErrorMessage: "can't Select Session: no session",
		},
		"Error with session ipdate": {
			session: core.Session{
				AccountID:    "Some-ID",
				Role:         "admin",
				RefreshToken: "RefreshToken-111",
				RequestHost:  "example.com",
				UserAgent:    "Some mozilla agent",
				ClientIP:     "127.0.0.1",
			},
			behaviorInsert: func(s *MockAccountStorage, session core.Session) {
				s.EXPECT().SelectSession(session).Return(core.Session{
					RefreshToken: session.RefreshToken,
					AccountID:    session.AccountID,
					Role:         session.Role,
					RequestHost:  session.RequestHost,
					UserAgent:    session.UserAgent,
					ClientIP:     session.ClientIP,
					Expired:      time.Now().Add(5 * time.Minute),
					Created:      time.Now().Add(-5 * time.Minute),
				}, nil)
			},
			behaviorRefresh: func(s *MockAccountStorage, session core.Session) {
				s.EXPECT().RefreshSession(gomock.Any()).Return(errors.New("error while update"))
			},
			expectedRefreshToken: "",
			expectedErrorMessage: "storege can't refress this session: error while update",
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			accountStorage := NewMockAccountStorage(ctrl)
			testCase.behaviorInsert(accountStorage, testCase.session)
			testCase.behaviorRefresh(accountStorage, testCase.session)

			accountService := AccountService{
				storage: accountStorage,
			}

			tokenPair, err := accountService.RefreshTokenpair(testCase.session)

			assert.Equal(t, testCase.expectedRefreshToken, tokenPair.RefreshToken)

			if err != nil {
				assert.Equal(t, testCase.expectedErrorMessage, err.Error())
			} else {
				token, _ := jwt.ParseWithClaims(tokenPair.AccessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, errInvalidSigningMethod
					}

					return []byte(signingKey), nil
				})

				claims, ok := token.Claims.(*Claims)
				if !ok {
					t.Error(errWrongTokenClaimType.Error())
				}

				assert.GreaterOrEqual(t, claims.ExpiresAt, claims.IssuedAt,
					"ExpiresAt cat't be less or equal to IssueAt")
				assert.Equal(t, testCase.session.AccountID, claims.Info.AccountID)
				assert.Equal(t, testCase.session.Role, claims.Info.Role)
				assert.Equal(t, testCase.session.RefreshToken, claims.Info.RefreshToken)
				assert.Equal(t, testCase.session.RequestHost, claims.Info.RequestHost)
				assert.Equal(t, testCase.session.UserAgent, claims.Info.UserAgent)
				assert.Equal(t, testCase.session.ClientIP, claims.Info.ClientIP)

			}
		})
	}
}

func TestService_logout(t *testing.T) {
	type mockBehavior func(s *MockAccountStorage, accountID string)

	testCasesTable := map[string]struct {
		accountID            string
		mockBehavior         mockBehavior
		expectedErrorMessage string
		wantError            bool
	}{
		"Succes": {
			accountID: "id-111",
			mockBehavior: func(s *MockAccountStorage, accountID string) {
				s.EXPECT().DeleteSesions(accountID).Return(nil)
			},
			expectedErrorMessage: "",
			wantError:            false,
		},
		"Should be an error": {
			accountID: "id-111",
			mockBehavior: func(s *MockAccountStorage, accountID string) {
				s.EXPECT().DeleteSesions(accountID).Return(errors.New("some error"))
			},
			expectedErrorMessage: "can't delete the account sessions: some error",
			wantError:            true,
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			// Init Deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			AccountStorage := NewMockAccountStorage(ctrl)
			testCase.mockBehavior(AccountStorage, testCase.accountID)

			accountService := AccountService{
				storage: AccountStorage,
			}

			err := accountService.Logout(testCase.accountID)
			if testCase.wantError {
				assert.Equal(t, testCase.expectedErrorMessage, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
