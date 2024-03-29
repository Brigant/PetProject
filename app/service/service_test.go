package service

import (
	"fmt"
	reflect "reflect"
	"testing"

	"github.com/Brigant/PetPorject/config"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	deps := Deps{
		AccountStorage:  NewMockAccountStorage(ctrl),
		DirectorStorage: NewMockDirectorStorage(ctrl),
		MovieStorage:    NewMockMovieStorage(ctrl),
		ListSorage:      NewMockListSorage(ctrl),
	}

	service := New(deps, config.Config{})

	sType := reflect.TypeOf(service)

	sValue := reflect.ValueOf(service)

	for i := 0; i < sType.NumField(); i++ {
		field := sType.Field(i)
		fieldVal := sValue.Field(i).Interface()

		fmt.Printf("%s: %v\n", field.Name, fieldVal)

		assert.NotEmpty(t, fieldVal, "All stucture field should be not nil")
	}
}
