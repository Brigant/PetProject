package core

type FavoriteList struct {
	ID       string `json:"id"`
	FilmID   string `json:"film_id"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
}

type WhishList struct {
	ID       string `json:"id"`
	FilmID   string `json:"film_id"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
}
