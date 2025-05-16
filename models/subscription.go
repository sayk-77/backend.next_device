package models

import "time"

type Subscription struct {
	Endpoint string `json:"endpoint"`
	Keys     struct {
		P256dh string `json:"p256dh"`
		Auth   string `json:"auth"`
	} `json:"keys"`
}

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Icon  string `json:"icon,omitempty"`
	Link  string `json:"link,omitempty"`
}

type PushSubscription struct {
	ID        uint   `gorm:"primaryKey"`
	Endpoint  string `gorm:"unique"`
	P256dh    string
	Auth      string
	UserID    int
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
}

type Keys struct {
	P256dh string `json:"p256dh"`
	Auth   string `json:"auth"`
}
