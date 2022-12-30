package models

import (
	"fmt"
	"log"

	"github.com/Akash-Parmar-0917/ecommerce/pkg/config"
	"github.com/go-errors/errors"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)


type Product struct {
	gorm.Model
	Name string `gorm:"" json:"name" form:"prodcutName"`
	Category Category `gorm:"foreignKey:CategoryId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"category"`
	// Category string `gorm:"" json:"category" form:"Category"`
	CategoryId uint
	SKU string `gorm:"" json:"sku" form:"sku"`
	Price int64 `gorm:"" json:"price" form:"price"`
	Specification datatypes.JSON `json:"specification" `
	Inventory Inventory `gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"inventory"`
}



func (p *Product) CreateProduct() *Product{
	db:=config.GetDB()
	db.Save(&p)
	log.Println("record crearted", *p)
	return p
}	


func GetAllProducts(pagination *PaginationData) (*[]Product, error) {
	var products []Product
	db:=config.GetDB()
	queryBuider := db.Limit(pagination.Limit).Offset(pagination.Offset).Order(pagination.Sort+",id").Model(&Product{}).Preload("Category").Preload("Inventory")
	var result error
	if(pagination.WhereQuery!=""){
		result = queryBuider.Where(pagination.WhereQuery).Find(&products).Error
	}else{
		result = queryBuider.Find(&products).Error
	}
	if result != nil {
		return nil, result
	}
	return &products, nil
}

// func GetAllProducts() []Product{
// 	var Products []Product
// 	db.Find(&Products)
// 	return Products
// }


func GetProductById(id int64) (*Product, *gorm.DB){
	var getProduct Product 
	db:=config.GetDB()
	db.Model(&Product{}).Where("id=?",id).Preload("Category").Preload("Inventory").Find(&getProduct)
	if(getProduct.ID==0){

		fmt.Println(errors.Errorf("not available id in produts table, id: %v ",id).ErrorStack())
	}
	return &getProduct,db
}

// func GetProdutsByIds(Id []int)(*[]Product,*gorm.DB){
// 	var getProduct []Product
// 	db:=db.Debug().Where("id IN",pq.Array(Id)).Preload("Category").Preload("Inventory").Find(&getProduct)
// 	return &getProduct,db
// }


func DeleteProduct(ID int64) Product{
	var product Product
	db:=config.GetDB()
	db.Where("id=?",ID).Unscoped().Delete(product)
	return product
}

