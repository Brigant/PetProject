package core

type Director struct {
	ID        string `json:"id"`
	Name      string `json:"name" binding:"required,min=2"`
	BirthDate string `json:"birth_date" binding:"required"`
	Created   string `json:"created"`
	Modified  string `json:"modified"`
}
