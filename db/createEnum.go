package db

import "gorm.io/gorm"

func createEnum(db *gorm.DB) {
	db.Exec("create type role as enum ('customer', 'admin', 'moderator')")
	db.Exec("create type payment_method as enum ('credit_card', 'paypal', 'bank_transfer')")
	db.Exec("create type payment_status as enum ('pending', 'completed', 'failed')")
	db.Exec("create type order_status as enum ('pending', 'paid', 'shipped', 'delivered', 'canceled', 'returned')")
}
