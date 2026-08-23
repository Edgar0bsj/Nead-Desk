package router

import (
	"nead-desk/src/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		// User
		v1.POST("/user", userHandler.Create)
		v1.GET("/user", userHandler.GetAll)
		v1.GET("/user/:id", userHandler.GetByID)
		v1.PUT("/user/:id", userHandler.Update)
		v1.DELETE("/user/:id", userHandler.Delete)
	}

	return r
}
