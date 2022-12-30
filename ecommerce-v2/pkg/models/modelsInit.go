package models

// import (
// 	"github.com/Akash-Parmar-0917/ecommerce/pkg/config"
// 	"gorm.io/gorm"
// )

// // func init(){
// // 	db=config.GetDB()

// // 	// db.Model(&Product{}).AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")
// // 	// db.Model(&Cart{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
// // 	// db.Model(&Inventory{}).AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")

// // 	// var products []Product
// // 	// db.Model(&Product{}).Where("category_id=1").Preload("Category").Preload("Inventory").First(&products)
// // 	// cart:=Cart{
// // 	// 	Products: products,
// // 	// 	Price: ,
// // 	// }
// // 	// var temp int64=4
// // 	// // fmt.Println("error while setting e1",err,v1)
// // 	// db.Model(&Cart{}).Save(&cart)
// // 	// tempCartProduct:=CartProduct{
// // 	// 	CartID: 1,
// // 	// 	ProductID: 2,
// // 	// 	Quantity: 3,
// // 	// 	E2: "should not create",
// // 	// }
// // 	// db.Model(&CartProduct{}).Create(&tempCartProduct)
// // 	// var cart Cart
// // 	// db.Model(&Cart{}).Preload("Product").First(&cart)
// // 	// fmt.Println(len(cart.Product))
// // 	// for _,product:=range cart.Product{
// // 	// 	db.Debug().Model(&CartProduct{}).Where("cart_id=?",cart.ID).Where("product_id=?",product.ID).Updates(map[string]interface{}{"e2":"fetched value"})
// // 	// }

// // 	// var cart1 Cart
// // 	// db.Debug().Model(&Cart{}).Preload("Products.Category").Preload("Products.Inventory").First(&cart1)
// // 	// fmt.Println(cart1.Products[0])
// // }