package models

type ProductVariants struct {
	Id           int    `json:"id" gorm:"unique;primary_key;AUTO_INCREMENT"`
	ProductId    int    `json:"productId"`
	VariantName  string `json:"variantName"`
	VariantValue string `json:"variantValue"`
}
