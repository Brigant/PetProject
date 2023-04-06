package core

import (
	"encoding/json"
	"errors"
	"time"
)

type Director struct {
	ID        string    `json:"id"  db:"id"`
	Name      string    `json:"name" binding:"required,min=2" db:"name"`
	BirthDate time.Time `json:"birth_date" binding:"required"  db:"birth_date"  `
	Created   string    `json:"created"  db:"created"`
	Modified  string    `json:"modified"  db:"modified"`
}

var ErrDublicatDirector = errors.New("there is the director with such data")

func (d *Director) UnmarshalJSON(data []byte) error {
	type Alias Director
	aux := &struct {
		BirthDate string `json:"birth_date"`
		*Alias
	}{
		Alias: (*Alias)(d),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02", aux.BirthDate)
	if err != nil {
		return err
	}

	d.BirthDate = t

	return nil
}

func (d *Director) MarshalJSON() ([]byte, error) {
	type Alias Director
	return json.Marshal(&struct {
		BirthDate string `json:"birth_date"`
		*Alias
	}{
		BirthDate: d.BirthDate.Format("2006-01-02"),
		Alias:     (*Alias)(d),
	})
}
