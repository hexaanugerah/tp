package models

type Hotel struct {
	ID          int
	Name        string
	City        string
	Description string
	Address     string
	Rating      float64
	PriceStart  int
	Image       string
	MapEmbedURL string
	Comments    []string
}
