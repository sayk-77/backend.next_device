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

func (s *UserService) SaveUserInfo(user *models.User) (string, error) {
	oldUserInfo, err := s.userRepo.GetUserByID(user.ID)
	if err != nil {
		return "", err
	}

	oldUserInfo.FirstName = user.FirstName
	oldUserInfo.LastName = user.LastName
	oldUserInfo.Email = user.Email

	return s.userRepo.UpdateUser(oldUserInfo)
}

func (s *UserService) ChangePassword(id uint, newPassword, oldPassword string) (string, error) {
	userPassword, err := s.userRepo.GetPasswordUserById(id)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(oldPassword)); err != nil {
		return "Не верный пароль", errors.New("invalid credentials")
	}

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return s.userRepo.ChangePassword(id, string(newHashedPassword))
}

func (s *UserService) AddNewAddress(id uint, address *models.Address) (string, error) {
	address.UserID = id
	return s.userRepo.AddNewAddress(address)
}

func (s *UserService) DeleteAddress(addressId, userId uint) (string, error) {
	return s.userRepo.DeleteAddress(addressId, userId)
}
