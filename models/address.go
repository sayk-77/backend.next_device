package models

type Address struct {
	Id         int    `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	City       string `json:"city"`
	Street     string `json:"street"`
	PostalCode string `json:"postCode"`
}
