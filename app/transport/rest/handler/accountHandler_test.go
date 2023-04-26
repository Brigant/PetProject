package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Brigant/PetPorject/app/core"
	"github.com/Brigant/PetPorject/logger"
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

func TestAccountHendler_login(t *testing.T) {
	log, err := logger.New("ERROR")
	if err != nil {
		t.FailNow()
	}

	type mockBehavior func(s *MockAccountService,
		phone, password string,
		c *gin.Context,
	)

	testCasesTable := map[string]struct {
		logger              *logger.Logger
		inputBody           string
		phone               string
		password            string
		session             core.Session
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		"Success": {
			logger:    log,
			inputBody: `{"phone":"+399999999","password":"password1234"}`,
			phone:     "+399999999",
			password:  "password1234",
			session:   core.Session{},
			mockBehavior: func(s *MockAccountService, phone, password string, c *gin.Context) {
				session := core.Session{
					RequestHost: c.Request.Host,
					UserAgent:   c.Request.UserAgent(),
					ClientIP:    c.ClientIP(),
				}
				s.EXPECT().Login(phone, password, session).Return(core.TokenPair{
					AccessToken:  "SomeAccesToken",
					RefreshToken: "SomeRefreshToken",
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"AccessToken":"SomeAccesToken","RefreshToken":"SomeRefreshToken"}`,
		},
		"Wrong request body": {
			logger:    log,
			inputBody: `{"phone":"+399999999","password":"password1234"`,
			phone:     "+399999999",
			password:  "password1234",
			session:   core.Session{},
			mockBehavior: func(s *MockAccountService, phone, password string, c *gin.Context) {
			},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"unexpected EOF"}`,
		},
		"Required password": {
			logger:    log,
			inputBody: `{"phone":"+399999999"}`,
			phone:     "+399999999",
			password:  "password1234",
			session:   core.Session{},
			mockBehavior: func(s *MockAccountService, phone, password string, c *gin.Context) {
			},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"Key: 'inputAccountData.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`,
		},
		"Required phone": {
			logger:    log,
			inputBody: `{"password":"password1234"}`,
			phone:     "+399999999",
			password:  "password1234",
			session:   core.Session{},
			mockBehavior: func(s *MockAccountService, phone, password string, c *gin.Context) {
			},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"Key: 'inputAccountData.Phone' Error:Field validation for 'Phone' failed on the 'required' tag"}`,
		},
		"Wrong phone": {
			logger:    log,
			inputBody: `{"phone":"399999999","password":"password1234"}`,
			phone:     "+399999999",
			password:  "password1234",
			session:   core.Session{},
			mockBehavior: func(s *MockAccountService, phone, password string, c *gin.Context) {
			},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"Key: 'inputAccountData.Phone' Error:Field validation for 'Phone' failed on the 'e164' tag"}`,
		},
		"Short password": {
			logger:    log,
			inputBody: `{"phone":"+399999999","password":"1234"}`,
			phone:     "+399999999",
			password:  "password1234",
			session:   core.Session{},
			mockBehavior: func(s *MockAccountService, phone, password string, c *gin.Context) {
			},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"Key: 'inputAccountData.Password' Error:Field validation for 'Password' failed on the 'min' tag"}`,
		},
		"Not ASCI password": {
			logger:    log,
			inputBody: `{"phone":"+399999999","password":"Парольsd1234"}`,
			phone:     "+399999999",
			password:  "password1234",
			session:   core.Session{},
			mockBehavior: func(s *MockAccountService, phone, password string, c *gin.Context) {
			},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"Key: 'inputAccountData.Password' Error:Field validation for 'Password' failed on the 'ascii' tag"}`,
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			// Init Deps.
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			accountService := NewMockAccountService(ctrl)

			accountHandler := AccountHandler{
				service: accountService,
				logger:  testCase.logger,
			}

			w := httptest.NewRecorder()

			c, r := gin.CreateTestContext(w)

			c.Request = httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte(testCase.inputBody)))

			testCase.mockBehavior(accountService,
				testCase.phone,
				testCase.password,
				c,
			)

			r.POST("/login", accountHandler.login)

			// Perform Request
			r.ServeHTTP(w, c.Request)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestAccountHendler_refreshToken(t *testing.T) {
	log, err := logger.New("ERROR")
	if err != nil {
		t.FailNow()
	}

	type mockBehavior func(s *MockAccountService, session core.Session)

	testCasesTable := map[string]struct {
		logger              *logger.Logger
		inputBody           string
		refreshToken        string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		"Success": {
			logger:       log,
			inputBody:    `{"RefreshToken": "fc182364-7122-4d4b-bd95-552b716224e2"}`,
			refreshToken: "fc182364-7122-4d4b-bd95-552b716224e2",
			mockBehavior: func(s *MockAccountService, session core.Session) {
				s.EXPECT().RefreshTokenpair(session).Return(
					core.TokenPair{
						AccessToken:  "SomeAccesToken",
						RefreshToken: "SomeRefreshToken",
					}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"AccessToken":"SomeAccesToken","RefreshToken":"SomeRefreshToken"}`,
		},
		"Bad refresh token": {
			logger:              log,
			inputBody:           `{"RefreshToken": "fc182364-7122-4d4b-bd95-552b716224e2"`,
			refreshToken:        "fc182364-7122-4d4b-bd95-552b716224e2",
			mockBehavior:        func(s *MockAccountService, session core.Session) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"unexpected EOF"}`,
		},
		"Bad UUID token": {
			logger:              log,
			inputBody:           `{"RefreshToken": "fc182364-7122-4d4b-bd95-552b716224e"}`,
			refreshToken:        "fc182364-7122-4d4b-bd95-552b716224e2",
			mockBehavior:        func(s *MockAccountService, session core.Session) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"invalid UUID length: 35"}`,
		},
		"Bad refreshToken key": {
			logger:              log,
			inputBody:           `{"RefreshTokensdf": "fc182364-7122-4d4b-bd95-552b716224e"}`,
			refreshToken:        "fc182364-7122-4d4b-bd95-552b716224e2",
			mockBehavior:        func(s *MockAccountService, session core.Session) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"invalid refresh token"}`,
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			accountService := NewMockAccountService(ctrl)

			accountHandler := AccountHandler{
				service: accountService,
				logger:  testCase.logger,
			}

			w := httptest.NewRecorder()

			c, r := gin.CreateTestContext(w)

			c.Request = httptest.NewRequest(http.MethodPost, "/refreshToken", bytes.NewBuffer([]byte(testCase.inputBody)))

			session := core.Session{
				RefreshToken: testCase.refreshToken,
				RequestHost:  c.Request.Host,
				UserAgent:    c.Request.UserAgent(),
				ClientIP:     c.ClientIP(),
				Expired:      time.Time{},
				Created:      time.Time{},
			}

			testCase.mockBehavior(accountService, session)

			r.POST("/refreshToken", accountHandler.refreshToken)

			// Perform Request
			r.ServeHTTP(w, c.Request)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestAccountHandler_logout(t *testing.T) {
	log, err := logger.New("ERROR")
	if err != nil {
		t.FailNow()
	}

	type mockBehavior func(s *MockAccountService)

	testCasesTable := map[string]struct {
		logger               *logger.Logger
		mockBehavior         mockBehavior
		ctxKey               string
		ctxVal               any
		expectedStatusCode   int
		expectedResponseBody string
	}{
		"Successful logout": {
			logger: log,
			mockBehavior: func(s *MockAccountService) {
				s.EXPECT().Logout(gomock.Any()).Return(nil)
			},
			ctxKey:               userCtx,
			ctxVal:               "accountID",
			expectedStatusCode:   200,
			expectedResponseBody: `{"action":"successful"}`,
		},
		"Bad string accountID": {
			logger: log,
			mockBehavior: func(s *MockAccountService) {
			},
			ctxKey:               userCtx,
			ctxVal:               111,
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"accountID is not string"}`,
		},
		"No contex in logout": {
			logger: log,
			mockBehavior: func(s *MockAccountService) {
			},
			ctxKey:               "",
			expectedStatusCode:   401,
			expectedResponseBody: `{"error":"unauthenticated"}`,
		},
		"Already logouted": {
			logger: log,
			mockBehavior: func(s *MockAccountService) {
				s.EXPECT().Logout(gomock.Any()).Return(core.ErrNoRowsEffected)
			},
			ctxKey:               userCtx,
			ctxVal:               "accountID",
			expectedStatusCode:   202,
			expectedResponseBody: `{"error":"no rows effected"}`,
		},
		"Some internal error": {
			logger: log,
			mockBehavior: func(s *MockAccountService) {
				s.EXPECT().Logout(gomock.Any()).Return(errors.New("some internal error"))
			},
			ctxKey:               userCtx,
			ctxVal:               "accountID",
			expectedStatusCode:   500,
			expectedResponseBody: `{"error":"some internal error"}`,
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			t.Run(name, func(t *testing.T) {
				gin.SetMode(gin.TestMode)

				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				accountService := NewMockAccountService(ctrl)
				testCase.mockBehavior(accountService)

				accountHandler := AccountHandler{
					service: accountService,
					logger:  testCase.logger,
				}

				w := httptest.NewRecorder()

				c, _ := gin.CreateTestContext(w)
				c.Set(testCase.ctxKey, testCase.ctxVal)

				accountHandler.logout(c)

				assert.Equal(t, testCase.expectedStatusCode, w.Code)
				assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
			})
		})
	}
}
