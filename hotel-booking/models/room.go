package models

type RoomType string

const (
	RoomTypeVIP      RoomType = "vip"
	RoomTypeDeluxe   RoomType = "deluxe"
	RoomTypeStandard RoomType = "biasa"
)

func (r RoomType) Label() string {
	switch r {
	case RoomTypeVIP:
		return "VIP"
	case RoomTypeDeluxe:
		return "Deluxe"
	default:
		return "Biasa"
	}
}

func (r RoomType) Code() string {
	switch r {
	case RoomTypeVIP:
		return "V"
	case RoomTypeDeluxe:
		return "D"
	default:
		return "B"
	}
}

type Room struct {
	ID           int
	HotelID      int
	RoomNumber   string
	Type         RoomType
	PricePerDay  int
	Capacity     int
	Beds         string
	Facilities   []string
	FloorMapURL  string
	HasBreakfast bool
	Available    bool
}
