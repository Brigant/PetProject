package core

import (
	"errors"
	"fmt"
	"strconv"
)

type Movie struct {
	ID          string `json:"id" db:"id"`
	Title       string `json:"title" binding:"required,min=1" db:"title"`
	Genre       string `json:"genre" binding:"required" db:"genre"`
	DirectorID  string `json:"director_id" binding:"required" db:"director_id"`
	Rate        int    `json:"rate" binding:"gte=0,lte=10" db:"rate"`
	ReleaseDate string `json:"release_date" binding:"required" db:"release_date"`
	Duration    int    `json:"duration" binding:"gte=1" db:"duration"`
	Created     string `json:"created" db:"created"`
	Modified    string `json:"modified" db:"modified"`
}

var (
	ErrUnallowedOffset      = errors.New("unallowed offset")
	ErrUnallowedFilterKey   = errors.New("unallowed filter key")
	ErrUnallowedSort        = errors.New("unallowed sort")
	ErrUnallowedLimit       = errors.New("unallowed limit")
	ErrForeignViolation     = errors.New("wrong foreign key")
	ErrUniqueMovie          = errors.New("dublicating the movie title with the such director")
	ErrNowMovieAdd          = errors.New("no movie added")
	ErrMovieNotFound        = errors.New("no movie with such ID")
	ErrUnallowedExportValue = errors.New("unallowed export value")
)

// QueryParams represent request query params
// that will be used on transport and repository level.
type QueryParams struct {
	Limit  string            `json:"limit"`
	Offset string            `json:"offset"`
	Filter map[string]string `json:"filter"`
	Sort   map[string]string `json:"sort"`
	Export string            `json:"export"`
}

// Set the default values to the Limit and Offset fields.
func (qp *QueryParams) SetDefaultValues() {
	if qp.Limit == "" {
		qp.Limit = "20"
	}

	if qp.Offset == "" {
		qp.Offset = "0"
	}
}

// Validate all fields of the query parameters.
func (qp QueryParams) Validate() error {
	var (
		minOffset          = 0
		maxOffset          = 1000
		allowedLimitVal    = []string{"20", "50", "100"}
		AllowedFilterKey   = []string{"genre", "rate"}
		AllowedSortKey     = []string{"rate", "release_date", "duration"}
		AllowedExportValue = []string{"csv", ""}
	)

	if notInSlice(qp.Limit, allowedLimitVal) {
		return ErrUnallowedLimit
	}

	offset, err := strconv.Atoi(qp.Offset)
	if err != nil {
		return fmt.Errorf("offset has to be an integer: %w", err)
	}

	if offset < minOffset || offset > maxOffset {
		return fmt.Errorf("offset should be in range from %v to %v: %w", minOffset, maxOffset, ErrUnallowedOffset)
	}

	for key, val := range qp.Filter {
		_ = val
		if notInSlice(key, AllowedFilterKey) {
			return ErrUnallowedFilterKey
		}

	}

	for key, val := range qp.Sort {
		_ = val
		if notInSlice(key, AllowedSortKey[:]) {
			return fmt.Errorf("key %v is bad: %w", key, ErrUnallowedSort)
		}

	}

	if notInSlice(qp.Export, AllowedExportValue) {
		return fmt.Errorf("value %v: %w", qp.Export, ErrUnallowedExportValue)
	}

	return nil
}

func notInSlice(element string, slice []string) bool {
	for _, s := range slice {
		if s == element {
			return false
		}
	}

	return true
}
