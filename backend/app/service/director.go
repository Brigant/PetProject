package service

import (
	"fmt"

	"github.com/Brigant/PetPorject/backend/app/core"
)

type DirectorService struct {
	storage DirectorStorage
}

func NewDirectorService(storage DirectorStorage) DirectorService {
	return DirectorService{storage: storage}
}

func (d DirectorService) CreatDirector(director core.Director) error {
	if err := d.storage.InsertDirector(director); err != nil {
		return fmt.Errorf("service get an error while InserDirector: %w", err)
	}

	return nil
}

func (d DirectorService) GetDirectorWithName(directorName string) (core.Director, error) {
	director, err := d.storage.SelectDirectorByName(directorName)
	if err != nil {
		return core.Director{}, fmt.Errorf("SelectDirectorByName returne error: %w", err)
	}

	return director, nil
}

func (d DirectorService) GetDirectorWithID(directorID string) (core.Director, error) {
	director, err := d.storage.SelectDirectorByID(directorID)
	if err != nil {
		return core.Director{}, fmt.Errorf("SelectDirectorByID returne error: %w", err)
	}

	return director, nil
}

func (d DirectorService) GetDirectorList() ([]core.Director, error) {
	directorsList, err := d.storage.SelectDirectorList()
	if err != nil {
		return nil, fmt.Errorf("SelectDirectorList returned the error: %w", err)
	}

	return directorsList, nil
}

func (d DirectorService) UpdateDirector() error {
	return nil
}
