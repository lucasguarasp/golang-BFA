package routes

import (
	"golang-bfa/internal/handlers"
	"golang-bfa/internal/middlewares"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App, healthHandler *handlers.HealthHandler, userHandler *handlers.UserHandler) {
	app.Use(middlewares.Logger())
	app.Get("/health", healthHandler.Health)
	app.Get("/users/:id", userHandler.GetUser)
	app.Post("/users", userHandler.CreateUser)
}