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

func (h *HotelHandler) HandleListHotels(c *fiber.Ctx) error {
	hotels, err := h.hotelStore.ListHotels(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(hotels)
}

func (h *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")

	hotel, err := h.hotelStore.GetHotel(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(hotel)
}
