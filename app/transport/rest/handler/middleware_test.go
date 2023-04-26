package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Brigant/PetPorject/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_UserIdentity(t *testing.T) {
	log, err := logger.New("INFO")
	if err != nil {
		t.Error("can't initialize logger")
	}

	type mockBehavior func(s *MockAccountService, accessToken string)

	tableTestCases := map[string]struct {
		logger               *logger.Logger
		mockBehavior         mockBehavior
		accessToken          string
		headerName           string
		headerValue          string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		"Success": {
			logger: log,
			mockBehavior: func(s *MockAccountService, accessToken string) {
				s.EXPECT().ParseToken(accessToken).Return("AccountID-111", "user", nil).Times(1)
			},
			accessToken:          "token",
			headerName:           authoriazahionHeader,
			headerValue:          "Bearer token",
			expectedStatusCode:   200,
			expectedResponseBody: "AccountID-111 user",
		},
		"Empty header": {
			logger:               log,
			mockBehavior:         func(s *MockAccountService, accessToken string) {},
			accessToken:          "token",
			headerName:           "",
			headerValue:          "Bearer token",
			expectedStatusCode:   401,
			expectedResponseBody: `{"error":"empty header, expecting Authorization header"}`,
		},
		"Invalid Bearer": {
			logger:               log,
			mockBehavior:         func(s *MockAccountService, accessToken string) {},
			accessToken:          "token",
			headerName:           authoriazahionHeader,
			headerValue:          "Bearerk token",
			expectedStatusCode:   401,
			expectedResponseBody: `{"error":"invalid header"}`,
		},
		"Empty token": {
			logger:               log,
			mockBehavior:         func(s *MockAccountService, accessToken string) {},
			accessToken:          "token",
			headerName:           authoriazahionHeader,
			headerValue:          "Bearerk ",
			expectedStatusCode:   401,
			expectedResponseBody: `{"error":"invalid header"}`,
		},
		"Service Failure": {
			logger: log,
			mockBehavior: func(s *MockAccountService, accessToken string) {
				s.EXPECT().ParseToken(accessToken).Return("", "", errors.New("failed to parse token")).Times(1)
			},
			accessToken:          "token",
			headerName:           authoriazahionHeader,
			headerValue:          "Bearer token",
			expectedStatusCode:   401,
			expectedResponseBody: `{"error":"failed to parse token"}`,
		},
	}

	for name, testCase := range tableTestCases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			accountService := NewMockAccountService(ctrl)
			testCase.mockBehavior(accountService, testCase.accessToken)

			mw := NewHandler(Deps{AccountService: accountService}, testCase.logger)

			// Build and setup Test Server
			gin.SetMode(gin.TestMode)
			r := gin.New()
			r.GET("/protected", mw.userIdentity, func(c *gin.Context) {
				accountID, _ := c.Get(userCtx)
				role, _ := c.Get(roleCtx)

				c.String(http.StatusOK, fmt.Sprintf("%s %s", accountID, role))
			})

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/protected", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)

			// Maker Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_AdminIdentitiy(t *testing.T) {
	log, err := logger.New("INFO")
	if err != nil {
		t.Error("can't initialize logger")
	}

	tableTestCases := map[string]struct {
		logger               *logger.Logger
		role                 string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		"Empty role": {
			logger:               log,
			role:                 "",
			expectedStatusCode:   401,
			expectedResponseBody: `{"error":"empty role"}`,
		},
		"Invalid role": {
			logger:               log,
			role:                 "user",
			expectedStatusCode:   401,
			expectedResponseBody: `{"error":"you are not admin"}`,
		},
	}

	for name, testCase := range tableTestCases {
		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			w := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(w)
			c.Set(roleCtx, testCase.role)

			middleware := Handler{
				Account:  AccountHandler{},
				Director: DirectorHandler{},
				Movie:    MovieHandler{},
				List:     ListHandler{},
				log:      testCase.logger,
			}

			// Invoke the middleware.
			middleware.adminIdentity(c)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
