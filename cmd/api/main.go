package main

import (
	"log"
	"nead-desk/src/handler"
	"nead-desk/src/router"
	"nead-desk/src/service"
	"nead-desk/src/storage"
)

func main() {

	userRepo := storage.NewUserMemoryStorage()

	userSvc := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userSvc)

	r := router.SetupRouter(userHandler)

	log.Println("Server running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
