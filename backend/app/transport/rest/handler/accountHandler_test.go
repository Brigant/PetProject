package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Brigant/PetPorject/backend/app/core"
	"github.com/Brigant/PetPorject/backend/logger"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAccountHandler_singUp(t *testing.T) {
	log, err := logger.New("INFO")
	if err != nil {
		t.FailNow()
	}

	type mockBehavior func(s *MockAccountService, account core.Account)

	testCasesTable := map[string]struct {
		logger              *logger.Logger
		inputBody           string
		account             core.Account
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		"Success": {
			logger:    log,
			inputBody: `{"phone":"+399999999","password":"password1234","age":15,"role":"admin"}`,
			account: core.Account{
				Phone:    "+399999999",
				Password: "password1234",
				Age:      15,
				Role:     "admin",
			},
			mockBehavior: func(s *MockAccountService, account core.Account) {
				s.EXPECT().CreateUser(account).Return("bla-bla-bla", nil)
			},
			expectedStatusCode:  201,
			expectedRequestBody: `{"userID":"bla-bla-bla"}`,
		},
		"Wrong Phone": {
			logger:              log,
			inputBody:           `{"phone":"399999999","password":"password1234","age":15,"role":"admin"}`,
			account:             core.Account{},
			mockBehavior:        func(s *MockAccountService, account core.Account) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"Key: 'Account.Phone' Error:Field validation for 'Phone' failed on the 'e164' tag"}`,
		},
		"To Short Password": {
			logger:              log,
			inputBody:           `{"phone":"+399999999","password":"1234567","age":15,"role":"admin"}`,
			account:             core.Account{},
			mockBehavior:        func(s *MockAccountService, account core.Account) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"Key: 'Account.Password' Error:Field validation for 'Password' failed on the 'min' tag"}`,
		},
		"To Long Password": {
			logger:              log,
			inputBody:           `{"phone":"+399999999","password":"123456790123456790123456790123456790123456790123456790123456790123456790123456790123456790123456790123456790123456790123456790123456790123456790123456790123456790123456790123456790123456790123456790123456790123456790123456790123456790123456790123456790123456790123456790","age":15,"role":"admin"}`,
			account:             core.Account{},
			mockBehavior:        func(s *MockAccountService, account core.Account) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"Key: 'Account.Password' Error:Field validation for 'Password' failed on the 'max' tag"}`,
		},
		"Not ASCII Password": {
			logger:              log,
			inputBody:           `{"phone":"+399999999","password":"1ігааіваі","age":15,"role":"admin"}`,
			account:             core.Account{},
			mockBehavior:        func(s *MockAccountService, account core.Account) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"Key: 'Account.Password' Error:Field validation for 'Password' failed on the 'ascii' tag"}`,
		},
		"Invalid character in the Password": {
			logger: log,
			inputBody: `{"phone":"+399999999","password":"sdfasdasd 
			sdfsdfs","age":15,"role":"admin"}`,
			account:             core.Account{},
			mockBehavior:        func(s *MockAccountService, account core.Account) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"invalid character '\\n' in string literal"}`,
		},
		"Age not int": {
			logger:              log,
			inputBody:           `{"phone":"+399999999","password":"123456789","age":"15","role":"admin"}`,
			account:             core.Account{},
			mockBehavior:        func(s *MockAccountService, account core.Account) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"json: cannot unmarshal string into Go struct field Account.age of type int"}`,
		},
		"Age to low": {
			logger:              log,
			inputBody:           `{"phone":"+399999999","password":"123456789","age":-30,"role":"admin"}`,
			account:             core.Account{},
			mockBehavior:        func(s *MockAccountService, account core.Account) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"Key: 'Account.Age' Error:Field validation for 'Age' failed on the 'gte' tag"}`,
		},
		"Age to high": {
			logger:              log,
			inputBody:           `{"phone":"+399999999","password":"123456789","age":130,"role":"admin"}`,
			account:             core.Account{},
			mockBehavior:        func(s *MockAccountService, account core.Account) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"Key: 'Account.Age' Error:Field validation for 'Age' failed on the 'lte' tag"}`,
		},
		"Role not in lowercase": {
			logger:              log,
			inputBody:           `{"phone":"+399999999","password":"123456789","age":30,"role":"Admin"}`,
			account:             core.Account{},
			mockBehavior:        func(s *MockAccountService, account core.Account) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"Key: 'Account.Role' Error:Field validation for 'Role' failed on the 'lowercase' tag"}`,
		},
		"Not available role": {
			logger:              log,
			inputBody:           `{"phone":"+399999999","password":"123456789","age":30,"role":"superuser"}`,
			account:             core.Account{},
			mockBehavior:        func(s *MockAccountService, account core.Account) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"Key: 'Account.Role' Error:Field validation for 'Role' failed on the 'checkRole' tag"}`,
		},
		"service Failure": {
			logger:    log,
			inputBody: `{"phone":"+399999999","password":"password1234","age":15,"role":"admin"}`,
			account: core.Account{
				Phone:    "+399999999",
				Password: "password1234",
				Age:      15,
				Role:     "admin",
			},
			mockBehavior: func(s *MockAccountService, account core.Account) {
				s.EXPECT().CreateUser(account).Return("", errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"error":"service failure"}`,
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			// Init Deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			accountService := NewMockAccountService(ctrl)
			testCase.mockBehavior(accountService, testCase.account)

			ah := AccountHandler{
				service: accountService,
				logger:  testCase.logger,
			}

			// Build and setup Test Server
			gin.SetMode(gin.TestMode)
			r := gin.New()

			if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
				if err := v.RegisterValidation("checkRole", checkRoleFunc); err != nil {
					ah.logger.Errorw("bind validator", "err", errValidatorBind.Error())
				}
			}

			r.POST("/", ah.singUp)

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testCase.inputBody))

			// Perform Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
