package repository

import (
	"errors"
	"gorm.io/gorm"
	"next_device/backend/models"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{db}
}

func (rr *ReviewRepository) CreateReview(review *models.Review) error {
	if result := rr.db.Create(&review); result.Error != nil {
		return result.Error
	}
	return nil
}

func (rr *ReviewRepository) GetAllReview() ([]*models.Review, error) {
	var reviews []*models.Review
	if result := rr.db.Find(&reviews); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return reviews, nil
		}
		return nil, result.Error
	}
	return reviews, nil
}

func (rr *ReviewRepository) GetReviewById(id uint) (*models.Review, error) {
	var review *models.Review
	if result := rr.db.First(&review, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return review, nil
}
func (rr *ReviewRepository) UpdateReview(review *models.Review) error {
	if result := rr.db.Save(&review); result.Error != nil {
		return result.Error
	}
	return nil
}

func (rr *ReviewRepository) DeleteReview(id uint) error {
	if result := rr.db.Delete(&models.Review{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}
