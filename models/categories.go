package models

type Categories struct {
	Id               int    `json:"id" gorm:"unique;primary_key;AUTO_INCREMENT"`
	Name             string `json:"name"`
	ParentCategoryId int    `json:"parentCategoryId"`
}
