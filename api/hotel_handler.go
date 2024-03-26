package api

import (
	"github.com/felipemagrassi/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
)

type HotelQueryParams struct {
	db.Pagination
	Rating int
}

type HotelHandler struct {
	roomStore  db.RoomStore
	hotelStore db.HotelStore
}

func NewHotelHandler(hotelStore db.HotelStore, roomStore db.RoomStore) *HotelHandler {
	return &HotelHandler{
		hotelStore: hotelStore,
		roomStore:  roomStore,
	}
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var params HotelQueryParams

	if err := c.QueryParser(&params); err != nil {
		return ErrBadRequest()
	}

}

func (h *HotelHandler) HandleGetHotel() {
	// TODO
}

func (h *HotelHandler) HandleCreateRoom() {
	// TODO
}
