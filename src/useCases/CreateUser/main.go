package usecases

import "github.com/gofiber/fiber/v2"

// CreateUser ...
func CreateUser(context *fiber.Ctx) error {
	var createUserController ICreateUserController = newCreateUserController()

	return createUserController.Handle(context)
}
