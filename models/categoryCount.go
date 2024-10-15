package models

type CategoryCount struct {
	Category Category `json:"category"`
	Count    int      `json:"count"`
}
