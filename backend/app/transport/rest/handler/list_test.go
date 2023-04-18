package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Brigant/PetPorject/backend/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestList_create(t *testing.T) {
	log, err := logger.New("DEBUG")
	if err != nil {
		t.FailNow()
	}

	type mockBehavior func(s *MockListsService)

	testCasesTable := map[string]struct {
		inputBody            string
		accountID            string
		userCtx              string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		"Successful case": {
			inputBody: `{"type":"favorite","movie_id":"6b823d5e-3d37-4617-a568-226e2e31a4f4"}`,
			accountID: "8c172d76-f750-4369-a5e2-27c877299168",
			userCtx:   userCtx,
			mockBehavior: func(s *MockListsService) {
				s.EXPECT().Create(gomock.Any()).Return("8c172d76-f750-4369-a5e2-27c877299168", nil).Times(1)
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `{"created with ID":"8c172d76-f750-4369-a5e2-27c877299168"}`,
		},
		"Wrong context": {
			inputBody:            `{"type":"favorite","movie_id":"6b823d5e-3d37-4617-a568-226e2e31a4f4"}`,
			accountID:            "8c172d76-f750-4369-a5e2-27c877299168",
			userCtx:              "wrong",
			mockBehavior:         func(s *MockListsService) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"no account found in contex"}`,
		},
		"Wrong accountID context": {
			inputBody:            `{"type":"favorite","movie_id":"6b823d5e-3d37-4617-a568-226e2e31a4f4"}`,
			accountID:            "",
			userCtx:              userCtx,
			mockBehavior:         func(s *MockListsService) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"invalid UUID length: 0"}`,
		},
		"Bad type": {
			inputBody:            `{"movie_id":"6b823d5e-3d37-4617-a568-226e2e31a4f4"}`,
			accountID:            "8c172d76-f750-4369-a5e2-27c877299168",
			userCtx:              userCtx,
			mockBehavior:         func(s *MockListsService) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"the movie list type should't be empty"}`,
		},
		"Bad movieID": {
			inputBody:            `{"type":"favorite"}`,
			accountID:            "8c172d76-f750-4369-a5e2-27c877299168",
			userCtx:              userCtx,
			mockBehavior:         func(s *MockListsService) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"the movie ID should't be empty"}`,
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			listService := NewMockListsService(ctrl)
			testCase.mockBehavior(listService)

			lh := NewListHandler(listService, log)

			response := httptest.NewRecorder()
			ctx, router := gin.CreateTestContext(response)
			router.POST("/list", lh.create)
			ctx.Request = httptest.NewRequest(http.MethodPost, "/list", strings.NewReader(testCase.inputBody))

			ctx.Set(testCase.userCtx, testCase.accountID)
			lh.create(ctx)

			assert.Equal(t, testCase.expectedStatusCode, response.Code)
			assert.Equal(t, testCase.expectedResponseBody, response.Body.String())
		})
	}
}
