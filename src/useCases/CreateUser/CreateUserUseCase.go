package usecases

import (
	"errors"
	"fmt"
	"time"

	"github.com/HallexCosta/clean-architecture-golang/src/entities"
	"golang.org/x/crypto/bcrypt"
)

// ICreateUserUseCase ...
type ICreateUserUseCase interface {
	execute(data *createUserRequestDTO)
	getError() error
}

// createUserUseCase ...
type createUserUseCase struct {
	error error
}

// newCreateUserUseCase ...
func newCreateUserUseCase() ICreateUserUseCase {
	var CreateUserUseCase ICreateUserUseCase = new(createUserUseCase)

	return CreateUserUseCase
}

// execute ...
func (createUserUseCase *createUserUseCase) execute(data *createUserRequestDTO) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)

	if err != nil {
		createUserUseCase.setError("Failed generate password hash")
	}

	var user *entities.User = &entities.User{
		ID:        1,
		Name:      data.Name,
		Email:     data.Email,
		Password:  string(passwordHash),
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
	}

	fmt.Println(user)

	return
}

// setError() ...
func (createUserUseCase *createUserUseCase) setError(message string) {
	createUserUseCase.error = errors.New(message)
}

// getError() ...
func (createUserUseCase *createUserUseCase) getError() error {
	if createUserUseCase.error == nil {
		createUserUseCase.error = nil
	}

	return createUserUseCase.error
}
