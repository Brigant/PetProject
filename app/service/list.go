package service

import (
	"fmt"

	"github.com/Brigant/PetPorject/app/core"
)

type ListService struct {
	storage ListSorage
}

func NewListService(storage ListSorage) ListService {
	return ListService{storage: storage}
}

func (s ListService) Create(list core.MovieList) (string, error) {
	listID, err := s.storage.Insert(list)
	if err != nil {
		return "", fmt.Errorf("create service got an error: %w", err)
	}

	return listID, nil
}

func (s ListService) GetAllAccountLists(condtitions []core.QuerySliceElement) ([]core.MovieList, error) {
	movieLists, err := s.storage.SelectAllUsersLists(condtitions)
	if err != nil {
		return nil, fmt.Errorf("select all users list got the error: %w", err)
	}

	return movieLists, nil
}

func (s ListService) AddMovieToList(listID, movieID string) error {
	if err := s.storage.InsertMovieToList(listID, movieID); err != nil {
		return fmt.Errorf("service add movie to list got error: %w", err)
	}

	return nil
}
