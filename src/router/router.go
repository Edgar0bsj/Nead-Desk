package router

import (
	"nead-desk/src/handler"
	"nead-desk/src/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.UserHandler, categoriHandler *handler.CategoriesHandler) *gin.Engine {
	r := gin.Default()

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", userHandler.UserAuthLogin)
		authRoutes.POST("/register", middleware.AuthRequired(), middleware.AuthAdmin(), userHandler.UserAuthRegister)
	}

	adminRoutes := r.Group("/admin", middleware.AuthRequired(), middleware.AuthAdmin())
	{
		adminRoutes.POST("/categories", categoriHandler.HandlerCreateCategories)
	}

	return r
}
