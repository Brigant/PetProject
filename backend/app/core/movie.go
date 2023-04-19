package core

import (
	"errors"
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
	minOffset               = 0
	maxOffset               = 1000
	minRate                 = 0
	maxRate                 = 10
	allowedLimitVal         = []string{"20", "50", "100"}
	allowedFilterKey        = []string{"genre", "rate"}
	allowedSortKey          = []string{"rate", "release_date", "duration"}
	allowedSortValue        = []string{"asc", "desc"}
	allowedExportValue      = []string{"csv", "none"}
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

type MovieCSV struct {
	Number       int      `csv:"Number"`
	Title        string   `csv:"Title" db:"title"`
	Genre        string   `csv:"Genre" db:"genre"`
	DirectorName string   `csv:"Director" db:"director_name"`
	Rate         int      `csv:"Rate" db:"rate"`
	ReleaseDate  DateTime `csv:"Release_Date" db:"release_date"`
	Duration     int      `csv:"Duration/Min" db:"duration"`
}
