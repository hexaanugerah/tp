package seed

import (
	"fmt"
	"hotel-booking/internal/models"
)

type citySeed struct {
	Name      string
	Landmark  string
	MapQuery  string
	PriceBase int
}

func SeedHotelsAndRooms() ([]models.Hotel, []models.Room) {
	cities := []citySeed{
		{Name: "Jakarta", Landmark: "Monas", MapQuery: "Monas Jakarta", PriceBase: 820000},
		{Name: "Bandung", Landmark: "Gedung Sate", MapQuery: "Gedung Sate Bandung", PriceBase: 700000},
		{Name: "Yogyakarta", Landmark: "Tugu Jogja", MapQuery: "Tugu Yogyakarta", PriceBase: 680000},
		{Name: "Bali", Landmark: "Tanah Lot", MapQuery: "Tanah Lot Bali", PriceBase: 930000},
		{Name: "Surabaya", Landmark: "Tugu Pahlawan", MapQuery: "Tugu Pahlawan Surabaya", PriceBase: 760000},
		{Name: "Medan", Landmark: "Istana Maimun", MapQuery: "Istana Maimun Medan", PriceBase: 740000},
		{Name: "Semarang", Landmark: "Lawang Sewu", MapQuery: "Lawang Sewu Semarang", PriceBase: 710000},
		{Name: "Malang", Landmark: "Bromo", MapQuery: "Bromo Malang", PriceBase: 690000},
		{Name: "Solo", Landmark: "Keraton Solo", MapQuery: "Keraton Solo", PriceBase: 660000},
		{Name: "Bogor", Landmark: "Kebun Raya", MapQuery: "Kebun Raya Bogor", PriceBase: 680000},
		{Name: "Makassar", Landmark: "Pantai Losari", MapQuery: "Pantai Losari Makassar", PriceBase: 720000},
		{Name: "Balikpapan", Landmark: "Pantai Kemala", MapQuery: "Balikpapan", PriceBase: 730000},
		{Name: "Lombok", Landmark: "Mandalika", MapQuery: "Mandalika Lombok", PriceBase: 820000},
		{Name: "Labuan Bajo", Landmark: "Pulau Komodo", MapQuery: "Labuan Bajo", PriceBase: 860000},
		{Name: "Palembang", Landmark: "Jembatan Ampera", MapQuery: "Jembatan Ampera Palembang", PriceBase: 680000},
		{Name: "Padang", Landmark: "Jam Gadang", MapQuery: "Padang", PriceBase: 670000},
		{Name: "Manado", Landmark: "Bunaken", MapQuery: "Bunaken Manado", PriceBase: 710000},
		{Name: "Pontianak", Landmark: "Khatulistiwa", MapQuery: "Tugu Khatulistiwa Pontianak", PriceBase: 660000},
		{Name: "Banjarmasin", Landmark: "Pasar Terapung", MapQuery: "Pasar Terapung Banjarmasin", PriceBase: 650000},
		{Name: "Jayapura", Landmark: "Teluk Youtefa", MapQuery: "Jayapura", PriceBase: 690000},
	}

	hotelPrefixes := []string{"Skyline", "Grand", "Royal", "Urban", "Heritage", "Signature", "Vista", "Sunrise", "Harmony", "Premier"}
	hotelSuffixes := []string{"Suites", "Residences", "Hotel", "Boutique", "Palace", "Inn", "Retreat", "Plaza", "Collection", "Haven"}

	hotels := make([]models.Hotel, 0, 200)
	hotelID := 1
	for _, city := range cities {
		for i := 0; i < 10; i++ {
			hotels = append(hotels, models.Hotel{
				ID:          hotelID,
				Name:        fmt.Sprintf("%s %s %s", city.Name, hotelPrefixes[i], hotelSuffixes[i]),
				City:        city.Name,
				Description: fmt.Sprintf("Hotel modern di %s dekat %s, cocok untuk bisnis dan liburan.", city.Name, city.Landmark),
				Address:     fmt.Sprintf("Pusat Kota %s", city.Name),
				Rating:      4.4 + (float64((i % 5)) * 0.1),
				PriceStart:  city.PriceBase + (i * 35000),
				Image:       fmt.Sprintf("https://picsum.photos/seed/hotel-%s-%d/640/420", city.Name, i+1),
				MapEmbedURL: fmt.Sprintf("https://maps.google.com/maps?q=%s&t=&z=13&ie=UTF8&iwloc=&output=embed", city.MapQuery),
				Comments: []string{
					"Lokasi strategis dekat pusat kota",
					"Pelayanan ramah dan check-in cepat",
				},
			})
			hotelID++
		}
	}

	roomTypes := []struct {
		Name       string
		Price      int
		Beds       int
		Capacity   int
		Facilities []string
	}{
		{"Biasa", 450000, 1, 2, []string{"WiFi", "AC", "TV 43 inch", "Kopi/Teh"}},
		{"Deluxe", 750000, 2, 3, []string{"WiFi cepat", "Smart TV", "Mini bar", "Bathtub"}},
		{"VIP", 1350000, 2, 4, []string{"Private lounge", "Jacuzzi", "Butler service", "City View"}},
	}

	rooms := make([]models.Room, 0, len(hotels)*30)
	roomID := 1
	for _, h := range hotels {
		for _, t := range roomTypes {
			for i := 1; i <= 10; i++ {
				rooms = append(rooms, models.Room{
					ID:            roomID,
					HotelID:       h.ID,
					Type:          t.Name,
					Name:          fmt.Sprintf("%s %s #%02d", h.City, t.Name, i),
					PricePerNight: t.Price + (i * 15000),
					Beds:          t.Beds,
					Capacity:      t.Capacity,
					Facilities:    t.Facilities,
					FloorPlanURL:  fmt.Sprintf("https://maps.google.com/maps?q=%s%%20hotel%%20floor%%20plan&t=&z=17&ie=UTF8&iwloc=&output=embed", h.City),
				})
				roomID++
			}
		}
	}

	return hotels, rooms
}
