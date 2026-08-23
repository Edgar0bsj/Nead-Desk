package router

import (
	"nead-desk/src/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.UserHandler, calledHandler *handler.CalledHandler) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		// User
		v1.POST("/user", userHandler.Create)
		v1.GET("/user", userHandler.GetAll)
		v1.GET("/user/:id", userHandler.GetByID)
		v1.PUT("/user/:id", userHandler.Update)
		v1.DELETE("/user/:id", userHandler.Delete)
		// Called
		v1.POST("/called", calledHandler.Create)
		v1.GET("/called", calledHandler.GetAll)
		v1.GET("/called/:id", calledHandler.GetByID)
		v1.PUT("/called/:id", calledHandler.Update)
		v1.DELETE("/called/:id", calledHandler.Delete)
	}

	return r
}
