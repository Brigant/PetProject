package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
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
	ErrMovieNotFound        = errors.New("no movie was found")
	ErrUnallowedExportValue = errors.New("unallowed export value")
	ErrUnallowedRateValue   = errors.New("unallowed rate value")
)

// QueryParams represent request query params
// that will be used on transport and repository level.
type QueryParams struct {
	Limit  string              `json:"limit"`
	Offset string              `json:"offset"`
	Filter []QuerySliceElement `json:"filter"`
	Sort   []QuerySliceElement `json:"sort"`
	Export string              `json:"export"`
}

type QuerySliceElement struct {
	Key string
	Val string
}

type MovieCSV struct {
	Number       int      `csv:"Number"`
	Title        string   `csv:"Title" db:"title"`
	Genre        string   `csv:"Genre" db:"genre"`
	DirectorName string   `csv:"Director" db:"director_name"`
	Rate         int      `csv:"Rate" db:"rate"`
	ReleaseDate  DateTime `csv:"Release_Date" db:"release_date"`
	Duration     int      `csv:"Duration/Min" db:"duration"`
}

type DateTime struct {
	time.Time
}

// Convert the internal date as CSV string
func (date *DateTime) MarshalCSV() (string, error) {
	return date.Time.Format("2006-01-02"), nil
}

func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("2006-01-02", csv)
	return err
}

func (date DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", date.Time.Format("2006-01-02"))), nil
}

func (date *DateTime) UnmarshalJSON(data []byte) error {
	var rawDate string

	err := json.Unmarshal(data, &rawDate)
	if err != nil {
		return fmt.Errorf("custom unmarshal error: %w", err)
	}

	parsedDate, err := time.Parse("2006-01-02", rawDate)
	if err != nil {
		return fmt.Errorf("date parsing error: %w", err)
	}

	*date = DateTime{parsedDate}

	return nil
}

// Set the default values to the Limit and Offset fields.
func (qp *QueryParams) SetDefaultValues() {
	if qp.Limit == "" {
		qp.Limit = "20"
	}

	if qp.Offset == "" {
		qp.Offset = "0"
	}

	if qp.Export == "" {
		qp.Export = "none"
	}
}

// Validate all fields of the query parameters.
func (qp QueryParams) Validate() error {
	var (
		minOffset          = 0
		maxOffset          = 1000
		minRate            = 0
		maxRate            = 10
		allowedLimitVal    = []string{"20", "50", "100"}
		allowedFilterKey   = []string{"genre", "rate"}
		allowedSortKey     = []string{"rate", "release_date", "duration"}
		allowedSortValue   = []string{"asc", "desc"}
		allowedExportValue = []string{"csv", "none"}
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
	
	for _, elem := range qp.Filter {
		if notInSlice(elem.Key, allowedFilterKey) {
			return ErrUnallowedFilterKey
		}

		if elem.Key == "rate" {
			i, err := strconv.Atoi(elem.Val)
			if err != nil {
				return fmt.Errorf("the value shlould be an integer: %w", err)
			}

			if i < minRate || i > maxRate {
				return fmt.Errorf("the value should be in range from 1 to 10 :%w", ErrUnallowedRateValue)
			}
		}
	}

	for _, elem := range qp.Sort {
		if notInSlice(elem.Key, allowedSortKey) {
			return fmt.Errorf("key %v is bad: %w", elem.Key, ErrUnallowedSort)
		}

		if notInSlice(elem.Val, allowedSortValue) {
			return fmt.Errorf("the sort value: %w", ErrUnallowedSort)
		}
	}

	if notInSlice(qp.Export, allowedExportValue) {
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
