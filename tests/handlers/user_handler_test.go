package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"strconv"
	"testing"
	"golang-bfa/internal/dto"
	"golang-bfa/internal/handlers"
	"golang-bfa/internal/repositories"
	"golang-bfa/internal/services"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_GetUser(t *testing.T) {
	app := fiber.New()
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	app.Get("/users/:id", userHandler.GetUser)

	// Criar um usuário
	user := userService.CreateUser("John", "john@example.com")
	req := httptest.NewRequest("GET", "/users/"+strconv.Itoa(user.ID), nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var response dto.UserResponse
	json.NewDecoder(resp.Body).Decode(&response)
	assert.Equal(t, user.ID, response.ID)
	assert.Equal(t, "John", response.Name)
}

func TestUserHandler_CreateUser(t *testing.T) {
	app := fiber.New()
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	app.Post("/users", userHandler.CreateUser)

	reqBody := dto.CreateUserRequest{Name: "Jane", Email: "jane@example.com"}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)

	var response dto.UserResponse
	json.NewDecoder(resp.Body).Decode(&response)
	assert.Equal(t, "Jane", response.Name)
	assert.Equal(t, "jane@example.com", response.Email)
}