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

func TestCreateUser(t *testing.T) {
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

func TestGenerateAccessToken(t *testing.T) {
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
				RequestHost:  "example.com",
				UserAgent:    "Some mozilla agent",
				ClientIP:     "127.0.0.1",
				ExpiredIn:    100,
				Created:      time.Now(),
			},
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			var accountService AccountService

			accessToken, _ := accountService.generateAccessToken(testCase.account, testCase.session)

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
