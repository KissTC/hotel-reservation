package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/kisstc/hotel-reservation/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":8000", "the listen address of the API server")
	flag.Parse()

	app := fiber.New()

	apiv1 := app.Group("api/v1")

	app.Get("/foo", handleFoo)

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*listenAddr)
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "working fine"})
}
