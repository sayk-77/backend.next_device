package models

type CategoryWithCountAndImage struct {
	Category      `json:"category"`
	Count         int64  `json:"count"`
	ImageCategory string `json:"images_category"`
}
