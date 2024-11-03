package service

import (
	"errors"
	"gorm.io/gorm"
	"next_device/backend/models"
	"next_device/backend/repository"
)

type CartService struct {
	cartRepo *repository.CartRepository
}

func NewCartService(cartRepo *repository.CartRepository) *CartService {
	return &CartService{cartRepo}
}

func (s *CartService) AddItem(userID uint, productID, quantity uint) error {
	var cart models.Cart
	cartId, err := s.cartRepo.GetCartByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cart.UserID = userID
			newCartID, err := s.cartRepo.CreateCart(&cart)
			if err != nil {
				return err
			}
			cartId = newCartID
		} else {
			return err
		}
	}

	cartItem := &models.CartItem{
		CartId:    cartId,
		ProductId: productID,
		Quantity:  quantity,
	}
	return s.cartRepo.AddItem(cartItem)
}

func (s *CartService) RemoveItem(userID, productID uint) error {
	return s.cartRepo.RemoveItem(userID, productID)
}

func (s *CartService) GetCartItems(userID uint) ([]models.CartItem, error) {
	return s.cartRepo.GetCartItems(userID)
}
