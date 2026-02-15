package database

import (
	"fmt"
	"sync"

	"hotel-booking/helpers"
	"hotel-booking/models"
)

type Database struct {
	mu            sync.RWMutex
	Users         []models.User
	Hotels        []models.Hotel
	Rooms         []models.Room
	Bookings      []models.Booking
	NextUserID    int
	NextBookingID int
}

var AppDB *Database

func Init() {
	AppDB = &Database{NextUserID: 3, NextBookingID: 1}
	seedUsers()
	seedHotelsAndRooms()
}

func seedUsers() {
	AppDB.Users = []models.User{
		{ID: 1, Name: "Admin", Email: "admin@hotel.com", Password: helpers.HashPassword("admin123"), Role: "admin"},
		{ID: 2, Name: "Guest", Email: "guest@hotel.com", Password: helpers.HashPassword("guest123"), Role: "user"},
	}
}

func seedHotelsAndRooms() {
	hotels := []models.Hotel{
		{ID: 1, Name: "Nusa Vista Resort", City: "Bali", Description: "Resort tropis dekat pantai dengan nuansa premium.", ImageURL: "https://images.unsplash.com/photo-1566073771259-6a8506099945", Rating: 9.2, PriceStart: 550000},
		{ID: 2, Name: "Bandung Sky Inn", City: "Bandung", Description: "Hotel modern pemandangan pegunungan.", ImageURL: "https://images.unsplash.com/photo-1551882547-ff40c63fe5fa", Rating: 8.9, PriceStart: 420000},
		{ID: 3, Name: "Jakarta Grand Central", City: "Jakarta", Description: "Lokasi strategis untuk bisnis dan liburan.", ImageURL: "https://images.unsplash.com/photo-1445019980597-93fa8acb246c", Rating: 9.0, PriceStart: 600000},
		{ID: 4, Name: "Yogyakarta Heritage Stay", City: "Yogyakarta", Description: "Konsep budaya Jawa dengan fasilitas lengkap.", ImageURL: "https://images.unsplash.com/photo-1455587734955-081b22074882", Rating: 8.7, PriceStart: 350000},
		{ID: 5, Name: "Lombok Ocean Breeze", City: "Lombok", Description: "Dekat spot snorkeling dan sunset point.", ImageURL: "https://images.unsplash.com/photo-1582719478250-c89cae4dc85b", Rating: 9.1, PriceStart: 480000},
		{ID: 6, Name: "Surabaya Urban Hotel", City: "Surabaya", Description: "Hotel stylish dekat pusat kuliner kota.", ImageURL: "https://images.unsplash.com/photo-1522708323590-d24dbb6b0267", Rating: 8.8, PriceStart: 390000},
		{ID: 7, Name: "Malang Green Valley", City: "Malang", Description: "Suasana sejuk cocok untuk keluarga.", ImageURL: "https://images.unsplash.com/photo-1520250497591-112f2f40a3f4", Rating: 8.6, PriceStart: 370000},
		{ID: 8, Name: "Semarang Port View", City: "Semarang", Description: "Akses mudah ke kawasan kota lama.", ImageURL: "https://images.unsplash.com/photo-1564501049412-61c2a3083791", Rating: 8.5, PriceStart: 360000},
		{ID: 9, Name: "Makassar Sunset Bay", City: "Makassar", Description: "Pemandangan laut dan kuliner seafood.", ImageURL: "https://images.unsplash.com/photo-1590490360182-c33d57733427", Rating: 8.8, PriceStart: 410000},
		{ID: 10, Name: "Medan City Lights", City: "Medan", Description: "Kamar luas untuk perjalanan bisnis.", ImageURL: "https://images.unsplash.com/photo-1578683010236-d716f9a3f461", Rating: 8.7, PriceStart: 430000},
	}
	AppDB.Hotels = hotels

	roomID := 1
	for _, h := range hotels {
		for i := 1; i <= 10; i++ {
			AppDB.Rooms = append(AppDB.Rooms, models.Room{ID: roomID, HotelID: h.ID, Name: fmt.Sprintf("Biasa %d", i), Type: "Biasa", Price: h.PriceStart, Stock: 1, Capacity: 2, Beds: 1, Facilities: roomFacilities("Biasa"), ImageURL: roomImage("Biasa", i)})
			roomID++
		}
		for i := 1; i <= 10; i++ {
			AppDB.Rooms = append(AppDB.Rooms, models.Room{ID: roomID, HotelID: h.ID, Name: fmt.Sprintf("Deluxe %d", i), Type: "Deluxe", Price: h.PriceStart + 250000, Stock: 1, Capacity: 3, Beds: 2, Facilities: roomFacilities("Deluxe"), ImageURL: roomImage("Deluxe", i)})
			roomID++
		}
		for i := 1; i <= 10; i++ {
			AppDB.Rooms = append(AppDB.Rooms, models.Room{ID: roomID, HotelID: h.ID, Name: fmt.Sprintf("VIP %d", i), Type: "VIP", Price: h.PriceStart + 500000, Stock: 1, Capacity: 4, Beds: 2, Facilities: roomFacilities("VIP"), ImageURL: roomImage("VIP", i)})
			roomID++
		}
	}
}

func roomFacilities(roomType string) []string {
	base := []string{"Wi-Fi cepat", "AC", "Smart TV", "Sarapan gratis"}
	switch roomType {
	case "VIP":
		return append(base, "Bathtub", "Airport transfer", "Executive lounge")
	case "Deluxe":
		return append(base, "Mini bar", "City view premium")
	default:
		return append(base, "Shower air hangat")
	}
}

func roomImage(roomType string, idx int) string {
	basic := []string{
		"https://images.unsplash.com/photo-1631049035182-249067d7618e",
		"https://images.unsplash.com/photo-1566665797739-1674de7a421a",
		"https://images.unsplash.com/photo-1618773928121-c32242e63f39",
	}
	deluxe := []string{
		"https://images.unsplash.com/photo-1590490360182-c33d57733427",
		"https://images.unsplash.com/photo-1582719508461-905c673771fd",
		"https://images.unsplash.com/photo-1578683010236-d716f9a3f461",
	}
	vip := []string{
		"https://images.unsplash.com/photo-1584132967334-10e028bd69f7",
		"https://images.unsplash.com/photo-1505693416388-ac5ce068fe85",
		"https://images.unsplash.com/photo-1591088398332-8a7791972843",
	}
	pick := func(items []string) string { return items[(idx-1)%len(items)] }
	switch roomType {
	case "VIP":
		return pick(vip)
	case "Deluxe":
		return pick(deluxe)
	default:
		return pick(basic)
	}
}

func (d *Database) Lock()   { d.mu.Lock() }
func (d *Database) Unlock() { d.mu.Unlock() }
func (d *Database) RLock()  { d.mu.RLock() }
func (d *Database) RUnlock() {
	d.mu.RUnlock()
}
