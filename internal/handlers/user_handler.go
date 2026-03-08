package handlers

import (
	"strconv"
	"golang-bfa/internal/dto"
	"golang-bfa/internal/services"
	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUser(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	user, err := h.service.GetUserByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	response := dto.ToUserResponse(user)
	return c.JSON(response)
}

func (h *UserHandler) CreateUser(c fiber.Ctx) error {
	var req dto.CreateUserRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}
	// Simples validação
	if req.Name == "" || req.Email == "" {
		return c.Status(400).JSON(fiber.Map{"error": "name and email required"})
	}
	user := h.service.CreateUser(req.Name, req.Email)
	response := dto.ToUserResponse(user)
	return c.Status(201).JSON(response)
}