package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kisstc/hotel-reservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Foo",
		LasttName: "Bar",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("james")
}
