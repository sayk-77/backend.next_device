package service

import (
	"next_device/backend/models"
	"next_device/backend/repository"
)

type ReviewService struct {
	reviewRep *repository.ReviewRepository
}

func NewReviewService(reviewRep *repository.ReviewRepository) *ReviewService {
	return &ReviewService{reviewRep}
}

func (s *ReviewService) CreateReview(review *models.Review) error {
	return s.reviewRep.CreateReview(review)
}

func (s *ReviewService) DeleteReview(reviewId uint) error {
	return s.reviewRep.DeleteReview(reviewId)
}

func (s *ReviewService) PublishReview(reviewId uint) error {
	return s.reviewRep.PublishReview(reviewId)
}

func (s *ReviewService) GetReviewById(reviewId uint) (*models.Review, error) {
	return s.reviewRep.GetReviewById(reviewId)
}

func (s *ReviewService) GetReviewForProduct(productId uint) ([]*models.Review, error) {
	return s.reviewRep.GetReviewForProduct(productId)
}

func (s *ReviewService) CreateReviewImages(reviewImages []models.ReviewImage) error {
	return s.reviewRep.CreateReviewImages(reviewImages)
}

func (s *ReviewService) GetAllReviews() ([]*models.Review, error) {
	return s.reviewRep.GetAllReview()
}

func (s *ReviewService) ChangeStatus(orderId uint, status string) error {
	return s.reviewRep.ChangeStatus(orderId, status)
}
