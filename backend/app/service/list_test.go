package service

import (
	"errors"
	"testing"

	"github.com/Brigant/PetPorject/backend/app/core"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestListService_Create(t *testing.T) {
	type mockBehavior func(s *MockListSorage, list core.MovieList)

	testCasesTable := map[string]struct {
		list                 core.MovieList
		mockBehavior         mockBehavior
		expectedErrorMessage string
		expectedID           string
		wantError            bool
	}{
		"Successful": {
			list: core.MovieList{},
			mockBehavior: func(s *MockListSorage, list core.MovieList) {
				s.EXPECT().Insert(list).Return("listID-111", nil).Times(1)
			},
			expectedID:           "listID-111",
			expectedErrorMessage: "",
			wantError:            false,
		},
		"Shoud be an error": {
			list: core.MovieList{},
			mockBehavior: func(s *MockListSorage, list core.MovieList) {
				s.EXPECT().Insert(list).Return("",
					errors.New("some error")).Times(1)
			},
			expectedID:           "",
			expectedErrorMessage: "create service got an error: some error",
			wantError:            true,
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			// Init Deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			listStorage := NewMockListSorage(ctrl)
			testCase.mockBehavior(listStorage, testCase.list)

			ls := ListService{
				storage: listStorage,
			}

			listID, err := ls.Create(testCase.list)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedErrorMessage,
					"We want get an error beceause the storage returned the error")
			} else {
				assert.NoError(t, err, "The error should be nil")
				assert.Equal(t, testCase.expectedID, listID)
			}
		})
	}
}

func TestListService_GetAllAccountLists(t *testing.T) {
	type mockBehavior func(s *MockListSorage)

	testCasesTable := map[string]struct {
		mockBehavior         mockBehavior
		expectedErrorMessage string
		expected             []core.MovieList
		wantError            bool
	}{
		"Successful": {
			mockBehavior: func(s *MockListSorage) {
				s.EXPECT().SelectAllUsersLists(gomock.Any()).Return([]core.MovieList{},
					nil).Times(1)
			},
			expected:  []core.MovieList{},
			wantError: false,
		},
		"Shoud be an error": {
			mockBehavior: func(s *MockListSorage) {
				s.EXPECT().SelectAllUsersLists(gomock.Any()).Return(
					nil, errors.New("some error")).Times(1)
			},
			expectedErrorMessage: "select all users list got the error: some error",
			wantError:            true,
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			// Init Deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			listStorage := NewMockListSorage(ctrl)
			testCase.mockBehavior(listStorage)

			ls := ListService{
				storage: listStorage,
			}

			_, err := ls.GetAllAccountLists([]core.QuerySliceElement{})

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedErrorMessage,
					"We want get an error beceause the storage returned the error")
			} else {
				assert.NoError(t, err, "The error should be nil")
			}
		})
	}
}

func TestListService_AddMovieToList(t *testing.T) {
	type mockBehavior func(s *MockListSorage, listID, movieID string)

	testCasesTable := map[string]struct {
		listID               string
		movieID              string
		mockBehavior         mockBehavior
		expectedErrorMessage string
		wantError            bool
	}{
		"Successful": {
			listID:  "listID-111",
			movieID: "movieID-222",
			mockBehavior: func(s *MockListSorage, listID, movieID string) {
				s.EXPECT().InsertMovieToList(listID, movieID).Return(nil).Times(1)
			},
			expectedErrorMessage: "",
			wantError:            false,
		},
		"Want error": {
			listID:  "listID-111",
			movieID: "movieID-222",
			mockBehavior: func(s *MockListSorage, listID, movieID string) {
				s.EXPECT().InsertMovieToList(listID, movieID).Return(
					errors.New("some error")).Times(1)
			},
			expectedErrorMessage: "service add movie to list got error: some error",
			wantError:            true,
		},
	}

	for name, testCase := range testCasesTable {
		t.Run(name, func(t *testing.T) {
			// Init Deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			listStorage := NewMockListSorage(ctrl)
			testCase.mockBehavior(listStorage, testCase.listID, testCase.movieID)

			ls := ListService{
				storage: listStorage,
			}

			err := ls.AddMovieToList(testCase.listID, testCase.movieID)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedErrorMessage,
					"We want get an error beceause the storage returned the error")
			} else {
				assert.NoError(t, err, "The error should be nil")
			}
		})
	}
}
