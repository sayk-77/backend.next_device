package models

type OrderItem struct {
	ID        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID   uint    `json:"orderId"`
	ProductID uint    `json:"productId"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	Price     float64 `gorm:"type:decimal(10,2);not null" json:"priceAtOrderTime"`

	Order Order `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
