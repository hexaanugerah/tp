package models

type Hotel struct {
	ID          int
	Name        string
	City        string
	Description string
	ImageURL    string
	Rating      float64
	PriceStart  int
}
