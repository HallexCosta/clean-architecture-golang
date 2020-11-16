package main

import (
	useCases "github.com/HallexCosta/clean-architecture-golang/src/useCases/CreateUser"
	"github.com/gofiber/fiber/v2"
)

// Routes ...
func Routes() {
	app.Post("/users", func(context *fiber.Ctx) error {
		return useCases.CreateUser(context)
	})
}
