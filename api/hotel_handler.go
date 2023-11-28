package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kisstc/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HoteHandler struct {
	store *db.Store
}

func NewHotelHandler(store *db.Store) *HoteHandler {
	return &HoteHandler{
		store: store,
	}
}

func (h *HoteHandler) HandleGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"hotelID": oid}
	rooms, err := h.store.Room.GetRooms(c.Context(), filter)
	if err != nil {
		return err
	}

	return c.JSON(rooms)
}

func (h *HoteHandler) HandleGetHotels(c *fiber.Ctx) error {
	hotels, err := h.store.Hotel.GetHotels(c.Context(), nil)
	if err != nil {
		return err
	}

	return c.JSON(hotels)
}

func (h *HoteHandler) HandleGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	// filter := bson.M{"hotelID": oid}
	hotel, err := h.store.Hotel.GetHotelByID(c.Context(), oid)
	if err != nil {
		return err
	}

	return c.JSON(hotel)
}
