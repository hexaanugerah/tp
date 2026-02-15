package controllers

import (
	"net/http"
	"sort"
	"strconv"

	"hotel-booking/database"
	"hotel-booking/models"
)

type RoomTypeOption struct {
	Type       string
	Price      int
	Capacity   int
	Beds       int
	Facilities []string
	ImageURL   string
	TotalRooms int
	RoomID     int
	HotelID    int
}

type RoomTypeGroup struct {
	Type  string
	Rooms []models.Room
}

func Home(w http.ResponseWriter, r *http.Request) {
	database.AppDB.RLock()
	hotels := database.AppDB.Hotels
	database.AppDB.RUnlock()
	render(w, "index.html", ViewData{"Request": r, "Hotels": hotels})
}

func HotelDetail(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	selectedType := r.URL.Query().Get("type")
	database.AppDB.RLock()
	defer database.AppDB.RUnlock()

	var selected *models.Hotel
	for _, h := range database.AppDB.Hotels {
		if h.ID == id {
			hotel := h
			selected = &hotel
			break
		}
	}
	if selected == nil {
		http.NotFound(w, r)
		return
	}

	optionsByType := map[string]RoomTypeOption{}
	roomsByType := map[string][]models.Room{}

	for _, room := range database.AppDB.Rooms {
		if room.HotelID != id {
			continue
		}

		roomsByType[room.Type] = append(roomsByType[room.Type], room)

		if existing, ok := optionsByType[room.Type]; ok {
			existing.TotalRooms++
			if room.Price < existing.Price {
				existing.Price = room.Price
				existing.RoomID = room.ID
				existing.Capacity = room.Capacity
				existing.Beds = room.Beds
				existing.Facilities = room.Facilities
				existing.ImageURL = room.ImageURL
			}
			optionsByType[room.Type] = existing
			continue
		}

		optionsByType[room.Type] = RoomTypeOption{
			Type:       room.Type,
			Price:      room.Price,
			Capacity:   room.Capacity,
			Beds:       room.Beds,
			Facilities: room.Facilities,
			ImageURL:   room.ImageURL,
			TotalRooms: 1,
			RoomID:     room.ID,
			HotelID:    room.HotelID,
		}
	}

	order := map[string]int{"Biasa": 1, "Deluxe": 2, "VIP": 3}
	var options []RoomTypeOption
	for _, item := range optionsByType {
		options = append(options, item)
	}
	sort.Slice(options, func(i, j int) bool {
		return order[options[i].Type] < order[options[j].Type]
	})

	var groups []RoomTypeGroup
	for _, roomType := range []string{"Biasa", "Deluxe", "VIP"} {
		rooms := roomsByType[roomType]
		if len(rooms) == 0 {
			continue
		}
		sort.Slice(rooms, func(i, j int) bool { return rooms[i].ID < rooms[j].ID })
		groups = append(groups, RoomTypeGroup{Type: roomType, Rooms: rooms})
	}

	if selectedType == "" && len(groups) > 0 {
		selectedType = groups[0].Type
	}

	var activeGroup RoomTypeGroup
	for _, group := range groups {
		if group.Type == selectedType {
			activeGroup = group
			break
		}
	}

	if activeGroup.Type == "" && len(groups) > 0 {
		activeGroup = groups[0]
		selectedType = groups[0].Type
	}

	render(w, "hotel_detail.html", ViewData{
		"Request":      r,
		"Hotel":        selected,
		"RoomOptions":  options,
		"RoomGroups":   groups,
		"ActiveGroup":  activeGroup,
		"SelectedType": selectedType,
	})
}
