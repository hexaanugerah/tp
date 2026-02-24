package services

import (
	"hotel-booking/internal/models"
	"hotel-booking/internal/repositories"
	"sort"
)

type HotelService struct {
	HotelRepo *repositories.HotelRepository
	RoomRepo  *repositories.RoomRepository
}

func (s *HotelService) Hotels() []models.Hotel { return s.HotelRepo.FindAll() }

func (s *HotelService) HotelsByCity(city string) []models.Hotel {
	return s.HotelRepo.FindByCity(city)
}

func (s *HotelService) Cities() []string {
	set := map[string]struct{}{}
	for _, h := range s.HotelRepo.FindAll() {
		set[h.City] = struct{}{}
	}
	out := make([]string, 0, len(set))
	for city := range set {
		out = append(out, city)
	}
	sort.Strings(out)
	return out
}

func (s *HotelService) HotelByID(id int) *models.Hotel { return s.HotelRepo.FindByID(id) }
func (s *HotelService) Rooms(hotelID int, roomType string) []models.Room {
	return s.RoomRepo.FindByHotelAndType(hotelID, roomType)
}
