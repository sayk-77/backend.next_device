package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Art           string  `json:"art"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	DiscountPrice float64 `json:"discountPrice"`
	Stock         int     `json:"stock"`
	CategoryId    int     `json:"categoryId"`
	BrandId       int     `json:"brandId"`
}
