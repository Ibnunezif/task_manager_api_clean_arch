package Usecases

import (
	"errors"
	"task-manager/Domain"
	"task-manager/Infrastructure"
	"task-manager/Repositories"
)

// UserUsecase defines user-related business logic
type UserUsecase interface {
	Register(username, password string) (*Domain.User, error)
	Login(username, password string) (string, error)
	Promote(username string) error
}

type userUsecase struct {
	repo Repositories.UserRepository
}

// Constructor
func NewUserUsecase(r Repositories.UserRepository) UserUsecase {
	return &userUsecase{repo: r}
}

// Register creates a new user
func (u *userUsecase) Register(username, password string) (*Domain.User, error) {
	// Hash password
	hashed, err := Infrastructure.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Create domain user with empty ID (Mongo will generate)
	user := Domain.NewUser("", username, hashed, "user")

	// Save user to DB
	return u.repo.Create(user)
}

// Login authenticates user and returns JWT
func (u *userUsecase) Login(username, password string) (string, error) {
	user, err := u.repo.GetByUsername(username)
	if err != nil {
		return "", err
	}

	// Check password
	if !Infrastructure.CheckPasswordHash(password, user.Password()) {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT
	token, err := Infrastructure.GenerateToken(user.ID(), user.Role())
	if err != nil {
		return "", err
	}

	return token, nil
}

// Promote user to admin
func (u *userUsecase) Promote(username string) error {
	return u.repo.UpdateRole(username, "admin")
}