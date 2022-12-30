package models

import (
	"gorm.io/gorm"
)

type Inventory struct{
	gorm.Model
	Quantity int64 `gorm:"" json:"quantity"`
	ProductName string`gorm:"" json:"product_name"`
	ProductId uint 
}
