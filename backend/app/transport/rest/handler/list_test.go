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
			inputBody: `{"type":"favorite"}`,
			accountID: "8c172d76-f750-4369-a5e2-27c877299168",
			userCtx:   userCtx,
			mockBehavior: func(s *MockListsService) {
				s.EXPECT().Create(gomock.Any()).Return("8c172d76-f750-4369-a5e2-27c877299168", nil).Times(1)
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `{"created with ID":"8c172d76-f750-4369-a5e2-27c877299168"}`,
		},
		"Wrong context": {
			inputBody:            `{"type":"favorite"}`,
			accountID:            "8c172d76-f750-4369-a5e2-27c877299168",
			userCtx:              "wrong",
			mockBehavior:         func(s *MockListsService) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"no account found in contex"}`,
		},
		"Wrong accountID context": {
			inputBody:            `{"type":"favorite"}`,
			accountID:            "",
			userCtx:              userCtx,
			mockBehavior:         func(s *MockListsService) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"invalid UUID length: 0"}`,
		},
		"Bad type": {
			inputBody:            `{}`,
			accountID:            "8c172d76-f750-4369-a5e2-27c877299168",
			userCtx:              userCtx,
			mockBehavior:         func(s *MockListsService) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"the movie list type should't be empty"}`,
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

func TestList_getAll(t *testing.T) {
	log, err := logger.New("DEBUG")
	if err != nil {
		t.FailNow()
	}

	type mockBehavior func(s *MockListsService, filter []core.QuerySliceElement)

	testCasesTable := map[string]struct {
		urlQuery             string
		accountID            string
		userCtx              string
		filter               []core.QuerySliceElement
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		"Successful case": {
			urlQuery:  `/?type=whish&type=favorite`,
			accountID: "8c172d76-f750-4369-a5e2-27c877299168",
			filter: []core.QuerySliceElement{
				{Key: "account_id", Val: "8c172d76-f750-4369-a5e2-27c877299168"},
				{Key: "type", Val: "whish"},
				{Key: "type", Val: "favorite"},
			},
			userCtx: userCtx,
			mockBehavior: func(s *MockListsService, filter []core.QuerySliceElement) {
				s.EXPECT().GetAllAccountLists(filter).Return([]core.MovieList{}, nil).Times(1)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `[]`,
		},
		"Error Unkown Condition": {
			urlQuery:  `/?type=whish&type=favorite`,
			accountID: "8c172d76-f750-4369-a5e2-27c877299168",
			filter: []core.QuerySliceElement{
				{Key: "account_id", Val: "8c172d76-f750-4369-a5e2-27c877299168"},
				{Key: "type", Val: "whish"},
				{Key: "type", Val: "favorite"},
			},
			userCtx: userCtx,
			mockBehavior: func(s *MockListsService, filter []core.QuerySliceElement) {
				s.EXPECT().GetAllAccountLists(filter).Return(nil,
					core.ErrUnkownConditionKey).Times(1)
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"condition has unknown parameters"}`,
		},
		"Unkown Error": {
			urlQuery:  `/?type=whish&type=favorite`,
			accountID: "8c172d76-f750-4369-a5e2-27c877299168",
			filter: []core.QuerySliceElement{
				{Key: "account_id", Val: "8c172d76-f750-4369-a5e2-27c877299168"},
				{Key: "type", Val: "whish"},
				{Key: "type", Val: "favorite"},
			},
			userCtx: userCtx,
			mockBehavior: func(s *MockListsService, filter []core.QuerySliceElement) {
				s.EXPECT().GetAllAccountLists(filter).Return(nil,
					errors.New("some error")).Times(1)
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"error":"some error"}`,
		},
		"No ID in context": {
			urlQuery:  `/?type=whish&type=favorite`,
			accountID: "8c172d76-f750-4369-a5e2-27c877299168",
			filter: []core.QuerySliceElement{
				{Key: "account_id", Val: "8c172d76-f750-4369-a5e2-27c877299168"},
				{Key: "type", Val: "whish"},
				{Key: "type", Val: "favorite"},
			},
			userCtx:              "bad",
			mockBehavior:         func(s *MockListsService, filter []core.QuerySliceElement) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"no account found in contex"}`,
		},
		"Mistake in query path": {
			urlQuery:  `/?typpe=whish&type=favorite`,
			accountID: "8c172d76-f750-4369-a5e2-27c877299168",
			filter: []core.QuerySliceElement{
				{Key: "account_id", Val: "8c172d76-f750-4369-a5e2-27c877299168"},
				{Key: "type", Val: "favorite"},
			},
			userCtx: userCtx,
			mockBehavior: func(s *MockListsService, filter []core.QuerySliceElement) {
				s.EXPECT().GetAllAccountLists(filter).Return([]core.MovieList{}, nil).Times(1)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `[]`,
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			listService := NewMockListsService(ctrl)
			testCase.mockBehavior(listService, testCase.filter)

			lh := NewListHandler(listService, log)

			response := httptest.NewRecorder()
			ctx, router := gin.CreateTestContext(response)
			router.POST("/list", lh.getAll)
			ctx.Request = httptest.NewRequest(http.MethodPost,
				"/list"+testCase.urlQuery, nil)

			ctx.Set(testCase.userCtx, testCase.accountID)
			lh.getAll(ctx)

			assert.Equal(t, testCase.expectedStatusCode, response.Code)
			assert.Equal(t, testCase.expectedResponseBody, response.Body.String())
		})
	}
}

func TestList_movieToList(t *testing.T) {
	log, err := logger.New("DEBUG")
	if err != nil {
		t.FailNow()
	}

	type mockBehavior func(s *MockListsService)

	testCasesTable := map[string]struct {
		inputBody            string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		"Successful case": {
			inputBody: `{"list_id":"e018e175-7813-4969-a99a-ed234afb2dd9","movie_id":"ca160814-59b3-4d1d-8bae-e3772fa0c6fb"}`,
			mockBehavior: func(s *MockListsService) {
				s.EXPECT().AddMovieToList("e018e175-7813-4969-a99a-ed234afb2dd9", "ca160814-59b3-4d1d-8bae-e3772fa0c6fb").Return(nil).Times(1)
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `{"action":"succesful"}`,
		},
		"Unique error": {
			inputBody: `{"list_id":"e018e175-7813-4969-a99a-ed234afb2dd9","movie_id":"ca160814-59b3-4d1d-8bae-e3772fa0c6fb"}`,
			mockBehavior: func(s *MockListsService) {
				s.EXPECT().AddMovieToList("e018e175-7813-4969-a99a-ed234afb2dd9", "ca160814-59b3-4d1d-8bae-e3772fa0c6fb").Return(
					core.ErrDuplicateRow).Times(1)
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"such record already exists"}`,
		},
		"Foreign key violation": {
			inputBody: `{"list_id":"e018e175-7813-4969-a99a-ed234afb2dd9","movie_id":"ca160814-59b3-4d1d-8bae-e3772fa0c6fb"}`,
			mockBehavior: func(s *MockListsService) {
				s.EXPECT().AddMovieToList("e018e175-7813-4969-a99a-ed234afb2dd9", "ca160814-59b3-4d1d-8bae-e3772fa0c6fb").Return(
					core.ErrForeignKeyViolation).Times(1)
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"some value has no reference to the list or to the movie"}`,
		},
		"Empty list_id": {
			inputBody: `{"movie_id":"ca160814-59b3-4d1d-8bae-e3772fa0c6fb"}`,
			mockBehavior: func(s *MockListsService) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"Key: 'requesMovieList.ListID' Error:Field validation for 'ListID' failed on the 'required' tag"}`,
		},
		"Empty movie_id": {
			inputBody: `{"movie_id":"ca160814-59b3-4d1d-8bae-e3772fa0c6fb"}`,
			mockBehavior: func(s *MockListsService) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"Key: 'requesMovieList.ListID' Error:Field validation for 'ListID' failed on the 'required' tag"}`,
		},
		"Wrong movieID": {
			inputBody: `{"list_id":"e018e175-7813-4969-a99a-ed234afb2dd9","movie_id":"a160814-59b3-4d1d-8bae-e3772fa0c6fb"}`,
			mockBehavior: func(s *MockListsService) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"invalid UUID length: 35"}`,
		},
		"Wrong listID": {
			inputBody: `{"list_id":"e018e175-7813-4969-a99a-ed234afb2dd","movie_id":"ca160814-59b3-4d1d-8bae-e3772fa0c6fb"}`,
			mockBehavior: func(s *MockListsService) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"invalid UUID length: 35"}`,
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

			router.POST("/list/add", lh.movieToList)

			ctx.Request = httptest.NewRequest(http.MethodPost, "/list/add", strings.NewReader(testCase.inputBody))

			router.ServeHTTP(response, ctx.Request)

			assert.Equal(t, testCase.expectedStatusCode, response.Code)
			assert.Equal(t, testCase.expectedResponseBody, response.Body.String())
		})
	}
}
