package main

import (
	"log"
	"nead-desk/src/handler"
	"nead-desk/src/router"
	"nead-desk/src/service"
	"nead-desk/src/storage"
)

func main() {

	// User
	userStorage := storage.NewUserMemoryStorage()
	userService := service.NewUserService(userStorage)
	userHandler := handler.NewUserHandler(userService)

	// Categories
	categoStorage := storage.NewCategoriesMemoryStorage()
	categoService := service.NewCategoriesService(categoStorage)
	categoHandler := handler.NewCategoriesHandler(categoService)

	r := router.SetupRouter(userHandler, categoHandler)

	log.Println("Server running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
