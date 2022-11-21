package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemID      string `sql:"type:VARCHAR(10)" gorm:"column:item_id"`
	ItemCode    string `sql:"type:VARCHAR(10)"`
	Description string `sql:"type:VARCHAR(100)" gorm:"column:description"`
	Quantity    int
	Price       int64 `gorm:"column:harga"`
	OrderItem   int
}
