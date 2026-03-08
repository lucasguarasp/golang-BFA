package repositories

import (
	"errors"
	"golang-bfa/internal/models"
)

type UserRepository struct {
	users  map[int]*models.User
	nextID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  make(map[int]*models.User),
		nextID: 1,
	}
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *UserRepository) Create(user *models.User) *models.User {
	user.ID = r.nextID
	r.nextID++
	r.users[user.ID] = user
	return user
}