package service

import (
	"fmt"

	"github.com/Brigant/PetPorject/backend/app/core"
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

func (s ListService) GetAllAccountLists(conditions core.ConditionParams) ([]core.MovieList, error) {
	movieLists, err := s.storage.SelectAllUsersLists(conditions)
	if err != nil {
		return nil, fmt.Errorf("select all users list got the error: %w", err)
	}

	return movieLists, nil
}
