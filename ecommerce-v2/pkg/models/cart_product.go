package models

import (
	"fmt"

	"github.com/Akash-Parmar-0917/ecommerce/pkg/config"
	"gorm.io/gorm"
)

type CartProduct struct {
	CartID uint `gorm:"primary_key"`
	// gorm.Model
	// Cart Cart `gorm:"foreignKey:CartID;references:ID"`
	ProductID uint `gorm:"primary_key"`
	// Product Product `gorm:"foreignKey:ProductID;references:ID"` 
	Quantity int64 `gorm:"default:1"`
}


func GetCartProduct(cartId int64,productId int64) *CartProduct{
	var cartProduct CartProduct
	db:=config.GetDB()
	db.Model(&CartProduct{}).Where("cart_id=?",cartId).Where("product_id=?",productId).Find(&cartProduct)
	return &cartProduct

}


func CheckProductForCart(productId int64,quantity int64,cartId int64) (*Product,error) {
	product,_:=GetProductById(productId)
	cartProduct:=GetCartProduct(cartId,productId)

	if(product.ID==0 || product.Inventory.Quantity<cartProduct.Quantity+quantity){
		return nil,fmt.Errorf("product_id %v doesn't exist or less than %v items in stock ",productId,quantity)
	}
	return product,nil
}

func AddProductToCart(cartID int64,productId int64,quantity int64)(*CartProduct,error){
	if _,err:=CheckProductForCart(productId,quantity,cartID,);(err!=nil){
		return  nil,err
	}
	db:=config.GetDB()
	cartProduct:=GetCartProduct(cartID,productId)
	cartProduct.CartID=uint(cartID)
	cartProduct.ProductID=uint(productId)
	cartProduct.Quantity+=quantity
	err:=db.Model(&CartProduct{}).Where("cart_id=?",cartID).Where("product_id=?",productId).Save(cartProduct).Error
	return cartProduct,err
}


func UpdateCart(cartID int64,productId int64,quantity int64) (*Cart,error){
	var cartProduct CartProduct
	cart,_:=GetCartById(cartID)
	if(cart.ID==0){
		return &Cart{},fmt.Errorf("cart with id %v not found",cartID)
	}
	product,_:=GetProductById(productId)
	// product,err:=CheckProductForCart(productId,quantity,cartID); if(err!=nil){
	// 	return &Cart{},err
	// }
	db:=config.GetDB()
	
	db.Model(&cartProduct).Where("cart_id=?",cartID).Where("product_id=?",productId).Find(&cartProduct)
	price:=cart.Price-cartProduct.Quantity*product.Price+quantity*product.Price
	cartProduct.Quantity=quantity
	db.Save(&cartProduct)
	db.Model(&Cart{}).Where("id=?",cartID).Update("price",price)
	db.Model(&Cart{}).Where("id=?",cartID).Preload("Product").Find(cart)
	return cart,nil
}

func (c *CartProduct) BeforeSave(db *gorm.DB) error {
	//
	fmt.Println("calling cart_products before save hook",c)
	// c.E2="default value"
	return nil
  }

  func (c *CartProduct) BeforeUpdate(db *gorm.DB) error {
	//
	fmt.Println("calling cart_products before update hook",c)
	return nil
}

func DeleteCartProduct (cartID int64,productID int64){
	db:=config.GetDB()
	var cartProduct CartProduct
	db.Model(&CartProduct{}).Where("cart_id=?",cartID).Where("product_id=?",productID).Find(&cartProduct)
	var cart Cart
	db.Model(&Cart{}).Where("id=?",cartID).Find(&cart)
	var product Product
	db.Model(&Product{}).Where("id=?",productID).Find(&product)
	cart.Price-=cartProduct.Quantity*product.Price
	db.Save(&cart)
	if(cartProduct.CartID!=0 && cartProduct.ProductID!=0){
		db.Delete(&cartProduct)
	}
	
}
