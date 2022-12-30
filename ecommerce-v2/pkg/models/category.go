package models

import (
	"fmt"
	"log"

	"github.com/Akash-Parmar-0917/ecommerce/pkg/config"
	"gorm.io/gorm"
)


type Category struct{
	gorm.Model
	CategoryName string `gorm:"" json:"category_name"`
}


func (c *Category)CreateCategory() *Category{
	var temp Category
	db:=config.GetDB()
	err:=db.Model(&Category{}).Where("category_name = ?",c.CategoryName).First(&temp).Error
	// fmt.Println("query result:",err.Error())
	if(err!=nil && err.Error()=="record not found"){
		db.Create(&c)
		log.Println("category record crearted ")
		return c;
	}
	return &temp;
}

func (c *Category)DeleteCategory() *Category{
	db:=config.GetDB()
	err:=db.Debug().Unscoped().Delete(c).Error
	if(err!=nil){
		fmt.Println("error while deleteing category",c.CategoryName)
		panic(err)
	}
	return c;
}

func GetCategoryByName(name string) (*Category,error){
	var category Category 
	db:=config.GetDB()
	err:=db.Where("category_name = ?",name).First(&category).Error
	return &category,err
}
func GetCategoryById(id uint) (*Category,error){
	var category Category 
	db:=config.GetDB()
	err:=db.Where("id = ?",id).First(&category).Error
	return &category,err
}

func GetAllCategories() (*[]Category, error) {
	var categories []Category
	db:=config.GetDB()
	db = db.Model(&Category{}).Find(&categories)
	if db.Error != nil {
		msg := db.Error
		return nil, msg
	}
	return &categories, nil
}