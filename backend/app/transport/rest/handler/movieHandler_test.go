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

func TestMovie_create(t *testing.T) {
	log, err := logger.New("INFO")
	if err != nil {
		t.FailNow()
	}

	type mockBehavior func(s *MockMovieService, movie core.Movie)

	testCasesTable := map[string]struct {
		inputBody            string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		"Successful case": {
			inputBody: `{
				"title":"Avatar2",
				"ganre":"Adventure",
				"director_id": "bed41cca-ee04-4975-ad7e-5b142e8a9306",
				"rate":1,
				"release_date":"2023-01-01",
				"duration":10800
			}`,
			mockBehavior: func(s *MockMovieService, movie core.Movie) {
				s.EXPECT().CreateMovie(gomock.Any()).Return(nil).Times(1)
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `{"action":"successful"}`,
		},
		"Bad json": {
			inputBody: `
				"titl":"Avatar2",
				"ganre":"Adventure",
				"director_id": "bed41cca-ee04-4975-ad7e-5b142e8a9306",
				"rate":1,
				"release_date":"2023-01-01",
				"duration":10800
			}`,
			mockBehavior: func(s *MockMovieService, movie core.Movie) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"json: cannot unmarshal string into Go value of type core.Movie"}`,
		},
		"Empty title": {
			inputBody: `{
				"title":"",
				"ganre":"Adventure",
				"director_id": "bed41cca-ee04-4975-ad7e-5b142e8a9306",
				"rate":1,
				"release_date":"2023-01-01",
				"duration":10800
			}`,
			mockBehavior: func(s *MockMovieService, movie core.Movie) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"Key: 'Movie.Title' Error:Field validation for 'Title' failed on the 'required' tag"}`,
		},
		"Empty ganre": {
			inputBody: `{
				"title":"Avatar2",
				"ganre":"",
				"director_id": "bed41cca-ee04-4975-ad7e-5b142e8a9306",
				"rate":1,
				"release_date":"2023-01-01",
				"duration":10800
			}`,
			mockBehavior: func(s *MockMovieService, movie core.Movie) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"Key: 'Movie.Ganre' Error:Field validation for 'Ganre' failed on the 'required' tag"}`,
		},
		"Wrong director ID": {
			inputBody: `{
				"title":"Avatar2",
				"ganre":"Adventure",
				"director_id": "wrong-uuid",
				"rate":1,
				"release_date":"2023-01-01",
				"duration":10800
			}`,
			mockBehavior: func(s *MockMovieService, movie core.Movie) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"invalid UUID length: 10"}`,
		},
		"To low rate": {
			inputBody: `{
				"title":"Avatar2",
				"ganre":"Adventure",
				"director_id": "bed41cca-ee04-4975-ad7e-5b142e8a9306",
				"rate":-1,
				"release_date":"2023-01-01",
				"duration":10800
			}`,
			mockBehavior: func(s *MockMovieService, movie core.Movie) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"Key: 'Movie.Rate' Error:Field validation for 'Rate' failed on the 'gte' tag"}`,
		},
		"To hight rate": {
			inputBody: `{
				"title":"Avatar2",
				"ganre":"Adventure",
				"director_id": "bed41cca-ee04-4975-ad7e-5b142e8a9306",
				"rate":11,
				"release_date":"2023-01-01",
				"duration":10800
			}`,
			mockBehavior: func(s *MockMovieService, movie core.Movie) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"Key: 'Movie.Rate' Error:Field validation for 'Rate' failed on the 'lte' tag"}`,
		},
		"To low duration": {
			inputBody: `{
				"title":"Avatar2",
				"ganre":"Adventure",
				"director_id": "bed41cca-ee04-4975-ad7e-5b142e8a9306",
				"rate":5,
				"release_date":"2023-01-01",
				"duration":0
			}`,
			mockBehavior: func(s *MockMovieService, movie core.Movie) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"Key: 'Movie.Duration' Error:Field validation for 'Duration' failed on the 'gte' tag"}`,
		},
		"Err Foreign Key Violation": {
			inputBody: `{
				"title":"Avatar2",
				"ganre":"Adventure",
				"director_id": "bed41cca-ee04-4975-ad7e-5b142e8a9306",
				"rate":1,
				"release_date":"2023-01-01",
				"duration":10800
			}`,
			mockBehavior: func(s *MockMovieService, movie core.Movie) {
				s.EXPECT().CreateMovie(gomock.Any()).Return(core.ErrForeignViolation).Times(1)
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"wrong foreign key"}`,
		},
		"Err Unique Movie": {
			inputBody: `{
				"title":"Avatar2",
				"ganre":"Adventure",
				"director_id": "bed41cca-ee04-4975-ad7e-5b142e8a9306",
				"rate":1,
				"release_date":"2023-01-01",
				"duration":10800
			}`,
			mockBehavior: func(s *MockMovieService, movie core.Movie) {
				s.EXPECT().CreateMovie(gomock.Any()).Return(core.ErrUniqueMovie).Times(1)
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"dublicating the movie title with the such director"}`,
		},
		"Internal Service Error": {
			inputBody: `{
				"title":"Avatar2",
				"ganre":"Adventure",
				"director_id": "bed41cca-ee04-4975-ad7e-5b142e8a9306",
				"rate":1,
				"release_date":"2023-01-01",
				"duration":10800
			}`,
			mockBehavior: func(s *MockMovieService, movie core.Movie) {
				s.EXPECT().CreateMovie(gomock.Any()).Return(errors.New("internal error")).Times(1)
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"error":"internal error"}`,
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			movie := core.Movie{}

			movieService := NewMockMovieService(ctrl)
			testCase.mockBehavior(movieService, movie)

			mh := NewMovieHandler(movieService, log)

			w := httptest.NewRecorder()

			c, r := gin.CreateTestContext(w)

			c.Request = httptest.NewRequest(http.MethodPost, "/movie", strings.NewReader(testCase.inputBody))

			r.POST("/movie", mh.create)

			r.ServeHTTP(w, c.Request)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestMovie_get(t *testing.T) {
	log, err := logger.New("INFO")
	if err != nil {
		t.FailNow()
	}

	type mockBehavior func(s *MockMovieService, movieID string)

	testCasesTable := map[string]struct {
		inputID              string
		paramName            string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		"Successful case": {
			inputID:   "6b823d5e-3d37-4617-a568-226e2e31a4f4",
			paramName: "id",
			mockBehavior: func(s *MockMovieService, movieID string) {
				s.EXPECT().Get(movieID).Return(core.Movie{
					ID:    "6b823d5e-3d37-4617-a568-226e2e31a4f4",
					Title: "TestTitle",
				}, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"id":"6b823d5e-3d37-4617-a568-226e2e31a4f4","title":"TestTitle","ganre":"","director_id":"","rate":0,"release_date":"","duration":0,"created":"","modified":""}`,
		},
		"Not found in params": {
			inputID:              "6b823d5e-3d37-4617-a568-226e2e31a4f4",
			paramName:            "wrong param name for test",
			mockBehavior:         func(s *MockMovieService, movieID string) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"No movieID param in path"}`,
		},
		"Wrong id": {
			inputID:              "6b823d5e-3d37-4617-a568-226e2e31a4f",
			paramName:            "id",
			mockBehavior:         func(s *MockMovieService, movieID string) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"invalid UUID length: 35"}`,
		},
		"Movie not found": {
			inputID:   "6b823d5e-3d37-4617-a568-226e2e31a4f4",
			paramName: "id",
			mockBehavior: func(s *MockMovieService, movieID string) {
				s.EXPECT().Get(movieID).Return(core.Movie{},
					core.ErrMovieNotFound)
			},
			expectedStatusCode:   http.StatusNotFound,
			expectedResponseBody: `{"error":"no movie with such ID"}`,
		},
		"Internal server error": {
			inputID:   "6b823d5e-3d37-4617-a568-226e2e31a4f4",
			paramName: "id",
			mockBehavior: func(s *MockMovieService, movieID string) {
				s.EXPECT().Get(movieID).Return(core.Movie{},
					errors.New("some error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"error":"some error"}`,
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			movieService := NewMockMovieService(ctrl)
			testCase.mockBehavior(movieService, testCase.inputID)

			mh := NewMovieHandler(movieService, log)

			w := httptest.NewRecorder()

			c, r := gin.CreateTestContext(w)

			c.Request = httptest.NewRequest(http.MethodGet, "/movie/"+testCase.inputID, nil)

			r.GET("/movie/:"+testCase.paramName, mh.get)

			r.ServeHTTP(w, c.Request)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
