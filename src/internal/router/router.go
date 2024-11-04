package router

import (
	"api-project/src/internal/controller"
	"api-project/src/internal/controller/middleware"
	"api-project/src/internal/repository"
	"api-project/src/internal/usecase"
	"api-project/src/pkg/db"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	db, err := db.OpenConnection()

	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUseCase)

	router.POST("/users", userController.CreateUser)
	router.POST("/login",userController.Login)
	router.PATCH("/users/:id", middleware.JwtAuthMiddleware(), userController.UpdateUser)
	router.DELETE("/users/:id", middleware.JwtAuthMiddleware(), userController.DeleteUser)

	return router
}
