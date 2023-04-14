package service

import (
	"errors"
	"testing"

	"github.com/Brigant/PetPorject/backend/app/core"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMovieService_create(t *testing.T) {
	type mockBehavior func(s *MockMovieStorage, movie core.Movie)

	testCasesTable := map[string]struct {
		movie                core.Movie
		mockBehavior         mockBehavior
		expectedErrorMessage string
		wantError            bool
	}{
		"Successful": {
			movie: core.Movie{},
			mockBehavior: func(s *MockMovieStorage, movie core.Movie) {
				s.EXPECT().InsertMovie(movie).Return(nil).Times(1)
			},
			expectedErrorMessage: "",
			wantError:            false,
		},
		"Wants error": {
			movie: core.Movie{
				Title: "Some title",
			},
			mockBehavior: func(s *MockMovieStorage, movie core.Movie) {
				s.EXPECT().InsertMovie(movie).Return(errors.New("some error")).Times(1)
			},
			expectedErrorMessage: "error happens while inserting movie: some error",
			wantError:            true,
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			// Init Deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mStorage := NewMockMovieStorage(ctrl)
			testCase.mockBehavior(mStorage, testCase.movie)

			ms := MovieService{
				movieStorage: mStorage,
			}

			err := ms.CreateMovie(testCase.movie)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedErrorMessage,
					"We want get an error beceause the storage returned the error")
			} else {
				assert.NoError(t, err, "The error should be nil")
			}
		})
	}
}

func TestMovieService_get(t *testing.T) {
	type mockBehavior func(s *MockMovieStorage, movieID string)

	testCasesTable := map[string]struct {
		movieID              string
		mockBehavior         mockBehavior
		expectedErrorMessage string
		wantError            bool
	}{
		"Successful": {
			movieID: "some-id-111",
			mockBehavior: func(s *MockMovieStorage, movieID string) {
				s.EXPECT().SelectMovieByID(movieID).Return(core.Movie{
					ID: movieID,
				}, nil)
			},
			expectedErrorMessage: "",
			wantError:            false,
		},
		"Wants error": {
			movieID: "some-id-111",
			mockBehavior: func(s *MockMovieStorage, movieID string) {
				s.EXPECT().SelectMovieByID(movieID).Return(core.Movie{},
					errors.New("some error"))
			},
			expectedErrorMessage: "service Get got the error: some error",
			wantError:            true,
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			// Init Deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mStorage := NewMockMovieStorage(ctrl)
			testCase.mockBehavior(mStorage, testCase.movieID)

			ms := MovieService{
				movieStorage: mStorage,
			}

			actualMovie, err := ms.Get(testCase.movieID)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedErrorMessage,
					"We want get an error beceause the storage returned the error")
			} else {
				assert.NoError(t, err, "The error should be nil")
				assert.Equal(t, testCase.movieID, actualMovie.ID)
			}
		})
	}
}
