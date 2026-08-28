package router

import (
	"nead-desk/src/handler"
	"nead-desk/src/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", userHandler.UserAuthLogin)
		authRoutes.POST("/register", middleware.AuthMiddleware(), middleware.AdminMiddle(), userHandler.UserAuthRegister)
	}

	return r
}
