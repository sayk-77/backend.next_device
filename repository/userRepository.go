package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"next_device/backend/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	if result := r.db.Find(models.User{}, &users); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	var existingUser models.User
	if err := r.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return fmt.Errorf("Пользователь с таким email уже существует")
	}

	if result := r.db.Create(user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user *models.User
	if result := r.db.First(&user, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	if result := r.db.Save(user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) DeleteUser(id uint) error {
	if result := r.db.Delete(&models.User{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("Почта или пароль не верны")
	}

	return &user, err
}
