package router

import (
	"nead-desk/src/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/auth/login", userHandler.UserAuth)

	// v1 := r.Group("/api/v1")
	// v1.Use(middleware.AuthMiddleware())
	// {

	// }

	return r
}
