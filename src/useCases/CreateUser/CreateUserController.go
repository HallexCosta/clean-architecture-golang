package usecases

import (
	"github.com/gofiber/fiber/v2"
)

// ICreateUserController ...
type ICreateUserController interface {
	Handle(context *fiber.Ctx) error
}

// createUserController ...
type createUserController struct{}

// newCreateUserController ...
func newCreateUserController() ICreateUserController {
	var CreateUserController ICreateUserController = new(createUserController)

	return CreateUserController
}

// Handle ...
func (createUserController *createUserController) Handle(context *fiber.Ctx) error {
	var data *createUserRequestDTO = new(createUserRequestDTO)

	context.BodyParser(data)

	var createUserUseCase ICreateUserUseCase = newCreateUserUseCase()

	createUserUseCase.execute(&createUserRequestDTO{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	})

	err := createUserUseCase.getError()

	if err != nil {
		return context.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return context.Status(201).JSON(fiber.Map{
		"message": "User created with successfully",
	})
}
