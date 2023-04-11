package service

import (
	"errors"
	"testing"
	"time"

	"github.com/Brigant/PetPorject/backend/app/core"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGreateDirector(t *testing.T) {
	type mockBehavior func(s *MockDirectorStorage, director core.Director)

	testCasesTable := map[string]struct {
		director             core.Director
		mockBehavior         mockBehavior
		expectedErrorMessage string
		wantError            bool
	}{
		"Successful": {
			director: core.Director{
				Name: "James Kameron",
				BirthDate: core.BirthDayType{
					Time: time.Date(2022, 12, 30, 0, 0, 0, 0, time.Local),
				},
			},
			mockBehavior: func(s *MockDirectorStorage, director core.Director) {
				s.EXPECT().InsertDirector(director).Return(nil)
			},
			wantError: false,
		},
		"Wants error": {
			director: core.Director{
				Name: "James Kameron",
				BirthDate: core.BirthDayType{
					Time: time.Date(2022, 12, 30, 0, 0, 0, 0, time.Local),
				},
			},
			mockBehavior: func(s *MockDirectorStorage, director core.Director) {
				s.EXPECT().InsertDirector(director).Return(
					errors.New("some storage error"),
				)
			},
			expectedErrorMessage: "service get an error while InserDirector: some storage error",
			wantError:            true,
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			// Init Deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			DirectorStorage := NewMockDirectorStorage(ctrl)
			testCase.mockBehavior(DirectorStorage, testCase.director)

			ds := DirectorService{
				storage: DirectorStorage,
			}

			err := ds.CreateDirector(testCase.director)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedErrorMessage,
					"We want get an error beceause the storage returned the error")
			} else {
				assert.NoError(t, err, "The error should be nil")
			}
		})
	}
}

func TestGetDirectorWithID(t *testing.T) {
	type mockBehavior func(s *MockDirectorStorage, directorID string)

	testCasesTable := map[string]struct {
		directorID           string
		mockBehavior         mockBehavior
		expectedDirector     core.Director
		expectedErrorMessage string
		wantError            bool
	}{
		"Successful": {
			directorID: "Some-ID-111",
			mockBehavior: func(s *MockDirectorStorage, directorID string) {
				s.EXPECT().SelectDirectorByID(directorID).Return(core.Director{
					ID:   directorID,
					Name: "James Kameron",
				}, nil)
			},
			expectedDirector: core.Director{
				ID:   "Some-ID-111",
				Name: "James Kameron",
			},
			wantError: false,
		},
		"Error": {
			directorID: "Some-ID-111",
			mockBehavior: func(s *MockDirectorStorage, directorID string) {
				s.EXPECT().SelectDirectorByID(directorID).Return(core.Director{},
					errors.New("some error"))
			},
			expectedDirector:     core.Director{},
			expectedErrorMessage: "selectDirectorByID returne error: some error",
			wantError:            true,
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			// Init Deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			DirectorStorage := NewMockDirectorStorage(ctrl)
			testCase.mockBehavior(DirectorStorage, testCase.directorID)

			ds := DirectorService{
				storage: DirectorStorage,
			}

			actualDirector, err := ds.GetDirectorWithID(testCase.directorID)
			if testCase.wantError {
				assert.Equal(t, testCase.expectedDirector, actualDirector, "The entety of director should has empty fields")
				assert.EqualError(t, err, testCase.expectedErrorMessage,
					"We want get an error beceause the storage returned the error")
			} else {
				assert.Equal(t, testCase.expectedDirector, actualDirector)
				assert.NoError(t, err, "The error should be nil")
			}
		})
	}
}

func TestGetDirectorList(t *testing.T) {
	type mockBehavior func(s *MockDirectorStorage)

	testCasesTable := map[string]struct {
		mockBehavior         mockBehavior
		expectedList         []core.Director
		expectedErrorMessage string
		wantError            bool
	}{
		"Successful": {
			mockBehavior: func(s *MockDirectorStorage) {
				s.EXPECT().SelectDirectorList().Return([]core.Director{
					{ID: "1"},
					{ID: "2"},
				}, nil)
			},
			expectedList: []core.Director{
				{ID: "1"},
				{ID: "2"},
			},
			wantError: false,
		},
		"Error": {
			mockBehavior: func(s *MockDirectorStorage) {
				s.EXPECT().SelectDirectorList().Return(nil,
					errors.New("some error"))
			},
			expectedList:         nil,
			expectedErrorMessage: "SelectDirectorList returned the error: some error",
			wantError:            true,
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			// Init Deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			DirectorStorage := NewMockDirectorStorage(ctrl)
			testCase.mockBehavior(DirectorStorage)

			ds := DirectorService{
				storage: DirectorStorage,
			}

			actualList, err := ds.GetDirectorList()
			if testCase.wantError {
				assert.Equal(t, testCase.expectedList, actualList, "The director list should be nil")
				assert.EqualError(t, err, testCase.expectedErrorMessage,
					"We want get an error beceause the storage returned the error")
			} else {
				assert.Equal(t, testCase.expectedList, actualList)
				assert.NoError(t, err, "The error should be nil")
			}
		})
	}
}
