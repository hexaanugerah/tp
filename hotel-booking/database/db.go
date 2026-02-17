package database

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"hotel-booking/models"
)

type Store struct {
	mu            sync.RWMutex
	Users         map[int]*models.User
	Hotels        map[int]*models.Hotel
	Rooms         map[int]*models.Room
	Bookings      map[int]*models.Booking
	NextUserID    int
	NextBookingID int
}

var DB *Store

func Init() {
	DB = &Store{
		Users:         map[int]*models.User{},
		Hotels:        map[int]*models.Hotel{},
		Rooms:         map[int]*models.Room{},
		Bookings:      map[int]*models.Booking{},
		NextUserID:    1,
		NextBookingID: 1,
	}
	seedData()
}

func seedData() {
	rand.Seed(time.Now().UnixNano())

	hotelNames := []string{
		"Aurora Sky Resort", "Nusa Breeze Suites", "Jakarta Central Grand", "Bandung Valley View",
		"Bali Ocean Crown", "Yogyakarta Heritage Stay", "Surabaya Prime Plaza", "Lombok Serenity Bay",
		"Bogor Highland Escape", "Malang Garden Hotel",
	}
	cities := []string{"Jakarta", "Bandung", "Bali", "Yogyakarta", "Surabaya", "Lombok", "Bogor", "Malang", "Semarang", "Medan"}
	desc := "Hotel modern dengan fasilitas lengkap: kolam renang, restoran premium, gym, wifi cepat, dan layanan 24 jam."
	amenities := []string{"WiFi Gratis", "Kolam Renang", "Sarapan", "Gym", "Spa", "Shuttle Bandara"}

	roomTypes := []models.RoomType{models.RoomTypeVIP, models.RoomTypeDeluxe, models.RoomTypeStandard}
	basePrice := map[models.RoomType]int{
		models.RoomTypeVIP:      1650000,
		models.RoomTypeDeluxe:   980000,
		models.RoomTypeStandard: 620000,
	}
	bedsByType := map[models.RoomType]string{
		models.RoomTypeVIP:      "2 King Bed",
		models.RoomTypeDeluxe:   "1 King Bed + 1 Single Bed",
		models.RoomTypeStandard: "1 Queen Bed",
	}
	facilitiesByType := map[models.RoomType][]string{
		models.RoomTypeVIP:      {"Private Lounge", "Bathtub", "Smart TV 65\"", "Mini Bar", "Nespresso"},
		models.RoomTypeDeluxe:   {"City View", "Work Desk", "Smart TV 50\"", "Rain Shower", "Sofa"},
		models.RoomTypeStandard: {"WiFi", "AC", "Smart TV 43\"", "Hot Shower", "Coffee & Tea"},
	}

	roomID := 1
	hotelID := 1
	for i := 0; i < 10; i++ {
		h := &models.Hotel{
			ID:             hotelID,
			Name:           hotelNames[i],
			City:           cities[i%len(cities)],
			Address:        fmt.Sprintf("Jl. Panorama No.%d", 10+i),
			Description:    desc,
			Rating:         4.2 + float64(rand.Intn(8))/10,
			ImageURL:       fmt.Sprintf("https://picsum.photos/seed/hotel%d/900/420", i+1),
			Amenities:      amenities,
			FloorMapURL:    fmt.Sprintf("https://picsum.photos/seed/floormap-hotel-%d/1200/500", i+1),
			LocationMapURL: fmt.Sprintf("https://maps.google.com/maps?q=%s&t=&z=13&ie=UTF8&iwloc=&output=embed", cities[i%len(cities)]),
		}

		for _, rt := range roomTypes {
			for n := 1; n <= 10; n++ {
				price := basePrice[rt] + rand.Intn(120000)
				r := &models.Room{
					ID:           roomID,
					HotelID:      hotelID,
					RoomNumber:   fmt.Sprintf("%s-%02d", rt.Code(), n),
					Type:         rt,
					PricePerDay:  price,
					Beds:         bedsByType[rt],
					Facilities:   facilitiesByType[rt],
					FloorMapURL:  fmt.Sprintf("https://picsum.photos/seed/room-map-%d-%s-%d/800/420", hotelID, rt, n),
					Capacity:     map[models.RoomType]int{models.RoomTypeVIP: 4, models.RoomTypeDeluxe: 3, models.RoomTypeStandard: 2}[rt],
					HasBreakfast: rt != models.RoomTypeStandard,
					Available:    true,
				}
				h.RoomIDs = append(h.RoomIDs, roomID)
				DB.Rooms[roomID] = r
				roomID++
			}
		}
		DB.Hotels[hotelID] = h
		hotelID++
	}

	admin := &models.User{ID: DB.NextUserID, Name: "Admin", Email: "admin@gostay.local", PasswordHash: "admin123", Role: models.RoleAdmin}
	DB.Users[admin.ID] = admin
	DB.NextUserID++
}
