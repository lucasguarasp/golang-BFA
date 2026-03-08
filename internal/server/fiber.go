package server

import (
	"golang-bfa/internal/config"
	"github.com/gofiber/fiber/v3"
)

func NewServer(cfg *config.Config) *fiber.App {
	app := fiber.New()
	return app
}