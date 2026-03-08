package services

import (
	"golang-bfa/internal/models"
	"golang-bfa/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) CreateUser(name, email string) *models.User {
	user := &models.User{
		Name:  name,
		Email: email,
	}
	return s.repo.Create(user)
}