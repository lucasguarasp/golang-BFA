package main

import (
	"log"
	"golang-bfa/internal/config"
	"golang-bfa/internal/handlers"
	"golang-bfa/internal/repositories"
	"golang-bfa/internal/routes"
	"golang-bfa/internal/server"
	"golang-bfa/internal/services"
	"golang-bfa/pkg/logger"
)

func main() {
	logger.Init()
	cfg := config.Load()
	app := server.NewServer(cfg)

	// Injeção de dependência
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	healthHandler := handlers.NewHealthHandler()

	routes.SetupRoutes(app, healthHandler, userHandler)

	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}