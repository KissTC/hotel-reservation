package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kisstc/hotel-reservation/db"
)

type HoteHandler struct {
	roomStore  db.RoomStore
	hotelStore db.HotelStore
}

func NewHotelHandler(hs db.HotelStore, rs db.RoomStore) *HoteHandler {
	return &HoteHandler{
		hotelStore: hs,
		roomStore:  rs,
	}
}

type HotelQueryParms struct {
	Rooms  bool
	Rating int
}

func (h *HoteHandler) HandleGetHotels(c *fiber.Ctx) error {
	var qparams HotelQueryParms
	err := c.QueryParser(&qparams)
	if err != nil {
		return err
	}

	fmt.Println(qparams)

	hotels, err := h.hotelStore.GetHotels(c.Context(), nil)
	if err != nil {
		return err
	}

	return c.JSON(hotels)
}
