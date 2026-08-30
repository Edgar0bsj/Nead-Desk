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
		// Authentication
		authRoutes.POST("/login", userHandler.UserAuthLogin)
		authRoutes.POST("/register", middleware.AuthRequired(), middleware.AuthAdmin(), userHandler.UserAuthRegister)
	}

	adminRoutes := r.Group("/admin", middleware.AuthRequired(), middleware.AuthAdmin())
	{
		// Categories
		adminRoutes.POST("/categories", categoriHandler.HandlerCreateCategories)
		adminRoutes.GET("/categories", categoriHandler.HandlerListAllCategores)
		adminRoutes.PATCH("/categories/:id", categoriHandler.HandlerUpdateCategores)
		adminRoutes.DELETE("/categories/:id", categoriHandler.HandlerDisableCategory)

		// Users
		adminRoutes.GET("/users", userHandler.HandlerFindAllUsers)
		adminRoutes.GET("/users/:id", userHandler.HandlerFindByIdUser)
		adminRoutes.PATCH("/users/:id", userHandler.HandlerChangeUser)
	}

	return r
}
