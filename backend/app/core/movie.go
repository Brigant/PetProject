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
	ErrForeignViolation = errors.New("wrong foreign key")
	ErrUniqueMovie      = errors.New("dublicating the movie title with the such director")
	ErrNowMovieAdd      = errors.New("no movie added")
	ErrNotFound         = errors.New("nothing was found")
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
