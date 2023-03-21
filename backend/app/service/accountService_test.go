package service

import (
	"errors"
	"testing"

	"github.com/Brigant/PetPorject/backend/app/core"
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
