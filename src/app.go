package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var app = App()

// App ...
func App() *fiber.App {
	return fiber.New()
}

// AppListen ...
func AppListen(p int) {
	port := fmt.Sprintf(":%d", p)
	app.Listen(port)
}
