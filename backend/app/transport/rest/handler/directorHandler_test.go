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

func TestDirector_create(t *testing.T) {
	log, err := logger.New("INFO")
	if err != nil {
		t.FailNow()
	}

	type mockBehavior func(s *MockDirectorService, director core.Director)

	testCasesTable := map[string]struct {
		logger              *logger.Logger
		inputBody           string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		"successfull case": {
			logger:    log,
			inputBody: `{"name":"Nois Perleone", "birth_date":"2022-12-30"}`,
			mockBehavior: func(s *MockDirectorService, director core.Director) {
				s.EXPECT().CreateDirector(gomock.Any()).Return(nil)
			},
			expectedStatusCode:  http.StatusCreated,
			expectedRequestBody: `{"action":"successful"}`,
		},
		"bad body request": {
			logger:    log,
			inputBody: `{"name":"Nois Perleone", "birth_date":"2022-12-30"`,
			mockBehavior: func(s *MockDirectorService, director core.Director) {
			},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"unexpected EOF"}`,
		},
		"No Name is body request": {
			logger:    log,
			inputBody: `{"name":"", "birth_date":"2022-12-30"}`,
			mockBehavior: func(s *MockDirectorService, director core.Director) {
			},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"Key: 'Director.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`,
		},
		"Bad BirthDay": {
			logger:    log,
			inputBody: `{"name":"Nois Perleone", "birth_date":"20222-12-30"}`,
			mockBehavior: func(s *MockDirectorService, director core.Director) {
			},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"parsing time \"20222-12-30\" as \"2006-01-02\": cannot parse \"2-12-30\" as \"-\""}`,
		},
		"service return error": {
			logger:    log,
			inputBody: `{"name":"Nois Perleone", "birth_date":"2022-12-30"}`,
			mockBehavior: func(s *MockDirectorService, director core.Director) {
				s.EXPECT().CreateDirector(gomock.Any()).Return(errors.New("some error"))
			},
			expectedStatusCode:  http.StatusInternalServerError,
			expectedRequestBody: `{"error":"some error"}`,
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			director := core.Director{}

			service := NewMockDirectorService(ctrl)
			testCase.mockBehavior(service, director)

			dh := NewDirectorHandler(service, testCase.logger)

			w := httptest.NewRecorder()
			c, r := gin.CreateTestContext(w)

			c.Request = httptest.NewRequest(http.MethodPost, "/director", strings.NewReader(testCase.inputBody))

			r.POST("director", dh.create)
			r.ServeHTTP(w, c.Request)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestDirector_getAll(t *testing.T) {
	log, err := logger.New("INFO")
	if err != nil {
		t.FailNow()
	}

	type mockBehavior func(s *MockDirectorService, id string)

	testCasesTable := map[string]struct {
		logger              *logger.Logger
		direcorID           string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		"successfull case": {
			logger:    log,
			direcorID: "dcabae88-1693-4349-af92-14704e4ffaab",
			mockBehavior: func(s *MockDirectorService, id string) {
				s.EXPECT().GetDirectorWithID(id).Return(core.Director{
					ID:        id,
					Name:      "James Kameron",
					BirthDate: core.BirthDayType{},
				}, nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `{"id":"dcabae88-1693-4349-af92-14704e4ffaab","name":"James Kameron","birth_date":"0001-01-01","created":"","modified":""}`,
		},
		"Internal error": {
			logger:    log,
			direcorID: "dcabae88-1693-4349-af92-14704e4ffaab",
			mockBehavior: func(s *MockDirectorService, id string) {
				s.EXPECT().GetDirectorWithID(id).Return(core.Director{},
					errors.New("some internal error"))
			},
			expectedStatusCode:  http.StatusInternalServerError,
			expectedRequestBody: `{"error":"some internal error"}`,
		},
		"Wrong id": {
			logger:              log,
			direcorID:           "Wrong-ID",
			mockBehavior:        func(s *MockDirectorService, id string) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"invalid UUID length: 8"}`,
		},
		"Empty id": {
			logger:              log,
			direcorID:           "",
			mockBehavior:        func(s *MockDirectorService, id string) {},
			expectedStatusCode:  http.StatusNotFound,
			expectedRequestBody: "404 page not found",
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewMockDirectorService(ctrl)
			testCase.mockBehavior(service, testCase.direcorID)

			dh := NewDirectorHandler(service, testCase.logger)

			w := httptest.NewRecorder()
			c, r := gin.CreateTestContext(w)

			path := "/director/" + testCase.direcorID

			c.Request = httptest.NewRequest(http.MethodGet, path, nil)

			r.GET("/director/:id", dh.get)
			r.ServeHTTP(w, c.Request)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestDirector_get(t *testing.T) {
	log, err := logger.New("INFO")
	if err != nil {
		t.FailNow()
	}

	type mockBehavior func(s *MockDirectorService)

	testCasesTable := map[string]struct {
		logger              *logger.Logger
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		"Successfull case": {
			logger: log,
			mockBehavior: func(s *MockDirectorService) {
				s.EXPECT().GetDirectorList().Return([]core.Director{
					{ID: "1"},
					{ID: "2"},
				}, nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `[{"id":"1","name":"","birth_date":"0001-01-01","created":"","modified":""},{"id":"2","name":"","birth_date":"0001-01-01","created":"","modified":""}]`,
		},
		"Internal error": {
			logger: log,
			mockBehavior: func(s *MockDirectorService) {
				s.EXPECT().GetDirectorList().Return([]core.Director{}, 
					errors.New("some internal error"))
			},
			expectedStatusCode:  http.StatusInternalServerError,
			expectedRequestBody: `{"error":"some internal error"}`,
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewMockDirectorService(ctrl)
			testCase.mockBehavior(service)

			dh := NewDirectorHandler(service, testCase.logger)

			w := httptest.NewRecorder()
			c, r := gin.CreateTestContext(w)

			c.Request = httptest.NewRequest(http.MethodGet, "/director", nil)

			r.GET("/director", dh.getAll)
			r.ServeHTTP(w, c.Request)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
