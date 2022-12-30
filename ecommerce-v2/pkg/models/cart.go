package models

import (
	"github.com/Akash-Parmar-0917/ecommerce/pkg/config"
	"gorm.io/gorm"
)

type Cart struct{
	// ID uint `gorm:"primaryKey, AUTO_INCREMENT"`
	gorm.Model
	Product []Product `gorm:"many2many:cart_products;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	// ProductId uint 
	// UserId uint
	Price int64 

}



func CreateCart(cart *Cart) (*Cart,*gorm.DB){
	db:=config.GetDB()
	db=db.Model(&Cart{}).Save(cart)
	return cart,db
}


func GetCartById(Id int64)(*Cart,*gorm.DB){
	var cart Cart
	db:=config.GetDB()
	db=db.Model(&Cart{}).Preload("Product.Category").Preload("Product.Inventory").Where("id=?",Id).First(&cart)
	return &cart,db
}


func DeleteCart(Id int64)(*Cart){
	var cart Cart
	db:=config.GetDB()
	db.Model(&Cart{}).Where("id=?",Id).Unscoped().Delete(&cart)
	return &cart
}