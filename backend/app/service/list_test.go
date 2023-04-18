package service

import (
	"errors"
	"testing"

	"github.com/Brigant/PetPorject/backend/app/core"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestListService_create(t *testing.T) {
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
