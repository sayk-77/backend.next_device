package models

type ProductWithMainImage struct {
	Id            uint    `json:"id"`
	Name          string  `json:"name"`
	SearchName    string  `json:"searchName"`
	Description   string  `json:"description"`
	CategoryTitle string  `json:"categoryTitle"`
	Image         string  `json:"image"`
	Price         float64 `json:"price"`
}
