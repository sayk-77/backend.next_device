package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"next_device/backend/models"
	"next_device/backend/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo}
}

func (s *UserService) Register(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)
	return s.userRepo.CreateUser(user)
}

func (s *UserService) Login(email, password string) (*models.User, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *UserService) GetUserById(id uint) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}
