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

// The sirvice with logic of creatinf of the director.
func (d DirectorService) CreateDirector(director core.Director) error {
	if err := d.storage.InsertDirector(director); err != nil {
		return fmt.Errorf("service get an error while InserDirector: %w", err)
	}

	return nil
}

// The service with logic of the getting of the one director.
func (d DirectorService) GetDirectorWithID(directorID string) (core.Director, error) {
	director, err := d.storage.SelectDirectorByID(directorID)
	if err != nil {
		return core.Director{}, fmt.Errorf("selectDirectorByID returne error: %w", err)
	}

	return director, nil
}

// The service returns the slice of the directors.
func (d DirectorService) GetDirectorList() ([]core.Director, error) {
	directorsList, err := d.storage.SelectDirectorList()
	if err != nil {
		return nil, fmt.Errorf("SelectDirectorList returned the error: %w", err)
	}

	return directorsList, nil
}
