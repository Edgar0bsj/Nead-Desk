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
	calledRepo := storage.NewCalledMemoryStorage()

	userSvc := service.NewUserService(userRepo)
	calledSvc := service.NewCalledService(calledRepo)

	userHandler := handler.NewUserHandler(userSvc)
	calledHandler := handler.NewCalledHandler(calledSvc)
	authHandler := handler.NewAuthHandler(userRepo)

	r := router.SetupRouter(userHandler, calledHandler, authHandler)

	log.Println("Server running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
