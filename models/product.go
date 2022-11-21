package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductID   string `sql:"unique;type:VARCHAR(68);not null" gorm:"column:product_id"`
	ProductName string `sql:"type:VARCHAR(68);not null"  gorm:"column:product_name"`
	Price       int64 `gorm:"column:price"`
	Description string `sql:"type:VARCHAR(100)" gorm:"column:description"`
}
