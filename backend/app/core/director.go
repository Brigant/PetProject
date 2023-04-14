package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type BirthDayType struct {
	time.Time
}

type Director struct {
	ID        string       `json:"id"  db:"id"`
	Name      string       `json:"name" binding:"required,min=2" db:"name"`
	BirthDate BirthDayType `json:"birth_date" binding:"required"  db:"birth_date"  `
	Created   string       `json:"created"  db:"created"`
	Modified  string       `json:"modified"  db:"modified"`
}

var (
	ErrDublicatDirector = errors.New("there is the director with such data")
	ErrNowDirectorAdded = errors.New("no director added")
	ErrNowDirectorFound = errors.New("no director found")
)

// Custom unmarshal function, for the custom type BrithDayType,
// which should format the time.Time field to "2006-01-02" patern.
func (b *BirthDayType) UnmarshalJSON(data []byte) error {
	var rawDate string

	err := json.Unmarshal(data, &rawDate)
	if err != nil {
		return fmt.Errorf("custom unmarshal error: %w", err)
	}

	parsedDate, err := time.Parse("2006-01-02", rawDate)
	if err != nil {
		return fmt.Errorf("date parsing error: %w", err)
	}

	*b = BirthDayType{parsedDate}

	return nil
}

// Custom marshal function, for the custom type BrithDayType,
// which should format the time.Time field to "2006-01-02" patern.
func (b BirthDayType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", b.Time.Format("2006-01-02"))), nil
}
