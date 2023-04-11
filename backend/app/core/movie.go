package core

import "errors"

type Movie struct {
	ID          string `json:"id" db:"id"`
	Title       string `json:"title" binding:"required,min=1" db:"title"`
	Ganre       string `json:"ganre" binding:"required" db:"ganre"`
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
	ErrNowMovieAdded    = errors.New("no movie added")
)
