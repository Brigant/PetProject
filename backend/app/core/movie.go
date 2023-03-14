package core

type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Ganre       string `json:"ganre"`
	DirectorID  string `json:"director_id"`
	Rate        int    `json:"rate"`
	ReleaseDate string `json:"release_date"`
	Duration    int    `json:"duration"`
	Created     string `json:"created"`
	Modified    string `json:"modified"`
}
