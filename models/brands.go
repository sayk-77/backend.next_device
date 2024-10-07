package models

type Brands struct {
	Id   int    `json:"id" gorm:"unique;primary_key;AUTO_INCREMENT"`
	Name string `json:"name"`
}
