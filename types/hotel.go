package types

type Hotel struct {
	ID       string   `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string   `json:"name" bson:"name"`
	Location string   `json:"location" bson:"location"`
	Rooms    []string `json:"rooms" bson:"rooms"`
}

type RoomType int

const (
	_ RoomType = iota
	SingleRoomType
	DoubleRoomType
	SeaSideRoomType
	DeluxeRoomType
)

type Room struct {
	ID        string   `json:"id,omitempty" bson:"_id,omitempty"`
	Type      RoomType `json:"type" bson:"type"`
	BasePrice float64  `json:"basePrice" bson:"basePrice"`
	Price     float64  `json:"price" bson:"price"`
	HotelID   string   `json:"hotelID" bson:"HotelID"`
}
