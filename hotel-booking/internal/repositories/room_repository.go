package repositories

import "hotel-booking/internal/models"

type RoomRepository struct{ Rooms []models.Room }

func (r *RoomRepository) FindByHotelAndType(hotelID int, roomType string) []models.Room {
	out := []models.Room{}
	for _, rm := range r.Rooms {
		if rm.HotelID == hotelID && (roomType == "" || rm.Type == roomType) {
			out = append(out, rm)
		}
	}
	return out
}

func (r *RoomRepository) FindByID(id int) *models.Room {
	for i := range r.Rooms {
		if r.Rooms[i].ID == id {
			return &r.Rooms[i]
		}
	}
	return nil
}
