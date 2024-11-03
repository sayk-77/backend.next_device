package models

type CartItem struct {
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id"`
	CartId    uint `json:"cartId"`
	Quantity  uint `json:"quantity"`
	ProductId uint `json:"productId"`
}
