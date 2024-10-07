package models

type ProductImages struct {
	Id        int    `json:"id" gorm:"unique;primary_key;AUTO_INCREMENT"`
	ProductId int    `json:"productId"`
	ImageUrl  string `json:"imageUrl"`
	IsMain    bool   `json:"isMain"`
}
