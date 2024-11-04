package controller

import (
	"api-project/src/internal/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createUserInput struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type loginInput struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type updateUserInput struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type deleteUserInput struct {
	Email string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type userController struct {
	UseCase usecase.UserUseCase
}

func NewUserController(useCase usecase.UserUseCase) *userController {
	return &userController{UseCase: useCase}
}

func (uc *userController) CreateUser(ctx *gin.Context) {

		var body createUserInput

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		user := &usecase.CreateUserInput{
			Name:      body.Name,
			Email:     body.Email,
			Password:  body.Password,
		}

		err := uc.UseCase.CreateUser(user)

		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (uc *userController) Login(ctx *gin.Context) {

	var body loginInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	authToken, err := uc.UseCase.Login(body.Email, body.Password)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"authToken": authToken})
}

func (uc *userController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var body updateUserInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		fmt.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := &usecase.UpdateUserInput{
		Name:      body.Name,
		Email:     body.Email,
		Password:  body.Password,
	}

	err := uc.UseCase.UpdateUser(id, user)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (uc *userController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var body deleteUserInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		fmt.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := &usecase.DeleteUserInput{
		Email:     body.Email,
		Password:  body.Password,
	}

	err := uc.UseCase.DeleteUser(id, user)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}