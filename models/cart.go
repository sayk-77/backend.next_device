package models

type Cart struct {
	ID     uint `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID uint `json:"userId"`

	Items []CartItem `gorm:"foreignKey:CartId" json:"items"`

	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
