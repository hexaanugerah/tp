package repositories

import "hotel-booking/internal/models"

type HotelRepository struct{ Hotels []models.Hotel }

func (r *HotelRepository) FindAll() []models.Hotel { return r.Hotels }

func (r *HotelRepository) FindByCity(city string) []models.Hotel {
	if city == "" {
		return r.Hotels
	}
	result := make([]models.Hotel, 0)
	for _, h := range r.Hotels {
		if h.City == city {
			result = append(result, h)
		}
	}
	return result
}

func (r *HotelRepository) FindByID(id int) *models.Hotel {
	for i := range r.Hotels {
		if r.Hotels[i].ID == id {
			return &r.Hotels[i]
		}
	}
	return nil
}
