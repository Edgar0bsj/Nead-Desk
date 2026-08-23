package router

import (
	"nead-desk/src/handler"
	"nead-desk/src/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.UserHandler, calledHandler *handler.CalledHandler, authHandler *handler.AuthHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/login", authHandler.Login)

	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthMiddleware())
	{
		admin := v1.Group("/admin")
		admin.Use(middleware.AdminMiddleware())
		{
			// User
			admin.POST("/user", userHandler.Create)
			admin.GET("/user", userHandler.GetAll)
			admin.GET("/user/:id", userHandler.GetByID)
			admin.PUT("/user/:id", userHandler.Update)
			admin.DELETE("/user/:id", userHandler.Delete)
		}

		// Called
		v1.POST("/called", calledHandler.Create)
		v1.GET("/called", calledHandler.GetAll)
		v1.GET("/called/:id", calledHandler.GetByID)
		v1.PUT("/called/:id", calledHandler.Update)
		v1.DELETE("/called/:id", calledHandler.Delete)
	}

	return r
}
