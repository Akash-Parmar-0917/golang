package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Akash-Parmar-0917/ecommerce/pkg/models"
	"github.com/gin-gonic/gin"
)



func AddCategory(c *gin.Context){
	
	addCategory:=&models.Category{}

    c.Bind(addCategory)
	// jsonData, _ := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(jsonData)
	// fmt.Println("addCategory",addCategory)
	addCategory=addCategory.CreateCategory()
	c.JSON(200,addCategory)
}

func DeleteCategory(c *gin.Context){
	categoryName:=c.Param("categoryName")
	category,err:=models.GetCategoryByName(categoryName)
	if(err==nil){
		category=category.DeleteCategory()
		c.JSON(200,category)
	}else{
		fmt.Println("error while finding category for DeleteCategory")
		c.JSON(404,gin.H{"status":"not found"})
	}
}

func GetCategories(c *gin.Context){
	cat,err:=models.GetAllCategories();
	if(err==nil){
		c.JSON(200,cat)
	}else{
		c.Status(http.StatusNotFound)
	}
}

func GetCategoryById(c *gin.Context){
	catId:=c.Param("categoryID");
	id,_:=strconv.Atoi(catId)
	category,err:=models.GetCategoryById(uint(id))
	if err==nil{
		c.JSON(200,category)
	}else{
		c.JSON(http.StatusNotFound,gin.H{"value":"record not found"})
	}

}

