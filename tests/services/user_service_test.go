package services_test

import (
	"testing"
	"golang-bfa/internal/repositories"
	"golang-bfa/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestUserService_GetUserByID(t *testing.T) {
	repo := repositories.NewUserRepository()
	service := services.NewUserService(repo)
	user := service.CreateUser("Test", "test@example.com")
	found, err := service.GetUserByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, found.ID)
	assert.Equal(t, "Test", found.Name)
}

func TestUserService_CreateUser(t *testing.T) {
	repo := repositories.NewUserRepository()
	service := services.NewUserService(repo)
	user := service.CreateUser("New", "new@example.com")
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "New", user.Name)
	assert.Equal(t, "new@example.com", user.Email)
}