package routers

import (
	"grace/controller"
	"grace/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	userService := &service.UserService{}
	userController := &controller.UserController{UserService: userService}
	junController := &controller.JunController{}

	// router.POST("/users", userController.CreateUser)
	router.GET("/users", userController.GetUsers)
	router.GET("/users/:id", userController.GetUser)
	router.PUT("/users/:id", userController.UpdateUser)
	// router.DELETE("/users/:id", userController.DeleteUser)

	router.GET("/api/get-header", junController.GetReqHeader)

	return router
}
