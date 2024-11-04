package usecase

import (
	"api-project/src/internal/auth"
	"api-project/src/internal/model"
	"api-project/src/internal/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Name      string
	Email     string
	Password  string
}

type UpdateUserInput struct {
	Name      string
	Email     string
	Password  string
}

type DeleteUserInput struct {
	Email string
	Password  string
}

type UserUseCase interface {
	CreateUser(user *CreateUserInput) error
	Login(email, password string) (string, error)
	UpdateUser(id string, user *UpdateUserInput) error
	DeleteUser(id string, user *DeleteUserInput) error
}

type userUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *userUseCase {
    return &userUseCase{repo: repo}
}

func (uc *userUseCase) CreateUser(user *CreateUserInput) error {

	userExists, err := uc.repo.GetUser(user.Email)

	if err != nil {
		return err
	}

	if userExists != nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashPassword)

	return uc.repo.CreateUser(&model.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
	})
}

func (uc *userUseCase) Login(email, password string) (string, error) {
	userExists, err := uc.repo.GetUser(email)

	if userExists == nil {
		return "",fmt.Errorf("user with email %s does not exists", email)
	}

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(password))

	if err != nil {
		return "", fmt.Errorf("invalid password")
	}

	token, err := auth.GenerateJwtToken(userExists.ID, password)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *userUseCase) UpdateUser(id string, user *UpdateUserInput) error {

	userExists, err := uc.repo.GetUser(user.Email)

	if userExists == nil {
		return fmt.Errorf("user with email %s does not exists", user.Email)
	}

	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(user.Password))

	if err != nil {
		return fmt.Errorf("invalid password")
	}

	return uc.repo.UpdateUser(id, &model.User{
		Name:      user.Name,
		Email:     user.Email,
	})
}

func (uc *userUseCase) DeleteUser(id string, user *DeleteUserInput) error {
	userExists, err := uc.repo.GetUserByID(id)

	if userExists == nil {
		return fmt.Errorf("user with id %s does not exists", id)
	}

	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(user.Password))

	if err != nil {
		return fmt.Errorf("invalid password")
	}

	return uc.repo.DeleteUser(id)
}