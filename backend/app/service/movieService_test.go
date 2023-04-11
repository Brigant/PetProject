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
				s.EXPECT().InsertMovie(movie).Return(nil)
			},
			expectedErrorMessage: "",
			wantError:            false,
		},
		"Wants error": {
			movie: core.Movie{},
			mockBehavior: func(s *MockMovieStorage, movie core.Movie) {
				s.EXPECT().InsertMovie(movie).Return(errors.New("some error"))
			},
			expectedErrorMessage: "some error",
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

			err := ms.movieStorage.InsertMovie(testCase.movie)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedErrorMessage,
					"We want get an error beceause the storage returned the error")
			} else {
				assert.NoError(t, err, "The error should be nil")
			}
		})
	}
}
