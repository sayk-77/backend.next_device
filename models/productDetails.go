package models

type ProductDetails struct {
	ID           uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID    uint   `json:"productId"`
	Processor    string `json:"processor"`
	GraphicsCard string `json:"graphicsCard"`
	RAM          string `json:"ram"`
	Storage      string `json:"storage"`
	Display      string `json:"display"`
	Camera       string `json:"camera"`
	Battery      string `json:"battery"`
	OS           string `json:"os"`
	Dimensions   string `json:"dimensions"`
	Weight       string `json:"weight"`
}
