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
	if result := r.db.Preload("Addresses").Preload("Orders").Find(&user, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) (string, error) {
	if result := r.db.Save(user); result.Error != nil {
		return "", result.Error
	}
	return "Данные успешно изменены", nil
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

func (r *UserRepository) ChangePassword(id uint, newPassword string) (string, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return "", err
	}
	user.PasswordHash = newPassword

	if err := r.db.Save(&user).Error; err != nil {
		return "", err
	}

	return "Пароль изменен", nil
}

func (r *UserRepository) GetPasswordUserById(id uint) (string, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return "", err
	}
	return user.PasswordHash, nil
}

func (r *UserRepository) AddNewAddress(address *models.Address) (string, error) {
	if err := r.db.Save(&address).Error; err != nil {
		return "", err
	}
	return "Адрес добавлен", nil
}

func (r *UserRepository) DeleteAddress(addressId, userId uint) (string, error) {
	result := r.db.Where("id = ? AND user_id = ?", addressId, userId).Delete(&models.Address{})

	if result.Error != nil {
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		return "", fmt.Errorf("Адрес не найден")
	}

	return "Адрес успешно удален", nil
}
