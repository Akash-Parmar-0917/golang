package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-errors/errors"

	"github.com/Akash-Parmar-0917/ecommerce/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)


func GetProducts (c *gin.Context){


	page := 1
	temp:=c.Param("page")
	temp=temp[1:]
	if(temp!=""){
		page,_=strconv.Atoi(temp)
	}
	sort := "category_id asc"
	// query :=  c.Request.URL.Query()
	decodedValue, err := url.QueryUnescape(c.Request.URL.String())
	index:=strings.Index(decodedValue,"?")
	decodedValue=decodedValue[index+1:]
	query,_:=url.ParseQuery(decodedValue)
	fmt.Println(query,decodedValue)
	whereQuery:=""
	
	// reqUrl:=tempstr[index+1:]
	var reqUrl string
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "sort":
			sort = queryValue
			if(len(reqUrl)>0){
				reqUrl+="&"
			}
			reqUrl+="sort="+queryValue
			break

		case "category_id":
			whereQuery="category_id="+queryValue
			if(len(reqUrl)>0){
				reqUrl+="&"
			}
			reqUrl+=whereQuery
		}
	}
	fmt.Println(reqUrl)
	pagination:=models.GetPaginationData(page,models.Product{},sort,whereQuery,reqUrl)
	productLists, err := models.GetAllProducts(&pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	// c.JSON(200,productLists)
	c.HTML(200,"index.html",gin.H{
		"products":productLists,
		"pagination":pagination,
	})
}


func AddProductPage(c *gin.Context){
	categoryList,err:=models.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.HTML(200,"addProduct.html",categoryList)
}

func UpdateProductPage(c *gin.Context){
	productID:=c.Query("productID")
	ID,err:=strconv.ParseInt(productID,0,0)
	if err!=nil{
		fmt.Println("error while Parsing in GetProductByID", err)
		return
	}
	productDetails,_:=models.GetProductById(ID)
	categoryList,err:=models.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	var temp map[string]interface{}
	json.Unmarshal(productDetails.Specification,&temp)
	fmt.Println(temp)
	c.HTML(200,"updateProduct.html",gin.H{
		"product":productDetails,
		"categories":categoryList,
		"specification": temp,
	})
}


func GetProductByID(c *gin.Context){
	productID:=c.Param("productID")
	ID,err:=strconv.ParseInt(productID,0,0)
	if err!=nil{
		fmt.Println(errors.Errorf("error while Parsing in GetProductByID: %w",err).ErrorStack())
		return
	}
	productDetails,_:=models.GetProductById(ID)
	if(productDetails.ID!=0){
		c.HTML(200,"productPage.html",productDetails)
	}else{
		c.JSON(http.StatusNotFound,gin.H{"value":"record not found"})
	}
	
}


func CreateProduct(c *gin.Context){
	
	CreateProduct:=&models.Product{}

    c.Bind(CreateProduct)
	var specification= make(map[string]string)
	key:=c.PostFormArray("key[]")
	value:=c.PostFormArray("value[]")
	categoryId,_:=strconv.Atoi(c.PostForm("category"))
	quantity,_:=strconv.ParseInt(c.PostForm("quantity"),0,0)
	sku:=c.PostForm("sku")
	price,_:=strconv.ParseInt(c.PostForm("price"),0,0)
	for i,k:= range key{
		specification[k]=value[i]
	}
	jsonString,err:=json.Marshal(specification)
	if err!=nil{
		fmt.Println(errors.Errorf("error while Parsing in controllers CreateProduct: %w",err).ErrorStack())
		return
	}
	CreateProduct.Specification=datatypes.JSON(jsonString)
	CreateProduct.Price=price
	CreateProduct.SKU=sku
	CreateProduct.Inventory=models.Inventory{
		Quantity: quantity,
	}

	cat,_:=models.GetCategoryById(uint(categoryId))
	CreateProduct.Category=*cat
	// utils.ParseBody(c.Request,CreateProduct)
	p:=CreateProduct.CreateProduct()
	c.JSON(200,p)
}


func DeleteProduct(c *gin.Context){
	productId:=c.Param("productID")
	ID,err:=strconv.ParseInt(productId,0,0)
	if err!=nil{
		fmt.Println(errors.Errorf("error while Parsing in controllers DeleteProduct: %w",err).ErrorStack())
		return
	}
	product:=models.DeleteProduct(ID)
	c.JSON(200,product)
}


func UpdateProduct(c *gin.Context){
	// updateProduct:=&models.Product{}

    // c.Bind(updateProduct)
	
	var tempSpec= make(map[string]string)
	key:=c.PostFormArray("key[]")
	value:=c.PostFormArray("value[]")
	quantity,_:=strconv.ParseInt(c.PostForm("quantity"),0,0)
	categoryid,_:=strconv.ParseInt(c.PostForm("category"),0,0)
	categoryId:=uint(categoryid)
	sku:=c.PostForm("sku")
	price,_:=strconv.ParseInt(c.PostForm("price"),0,0)
	name:=c.PostForm("prodcutName")
	for i,k:= range key{
		if(k==""){
			continue
		}
		tempSpec[k]=value[i]
	}
	jsonString,err:=json.Marshal(tempSpec)
	if err!=nil{
		fmt.Println(errors.Errorf("error while Parsing in controllers UpdateProduct: %w",err).ErrorStack())
		return
	}
	specification:=datatypes.JSON(jsonString)

	// var updateProduct = &models.Product{}
	// utils.ParseBody(c.Request,updateProduct)
	productId:=c.Param("productID")
	ID,err:=strconv.ParseInt(productId,0,0)
	if err!=nil{
		fmt.Println(errors.Errorf("error while Parsing in controllers UpdateProduct: %w",err).ErrorStack())
		return
	}
	productDetails,db:= models.GetProductById(ID)
	if name!=""{
		productDetails.Name=name
	}
	if categoryId!=productDetails.CategoryId{
		cat,_:=models.GetCategoryById(categoryId)
		productDetails.Category=*cat
		// productDetails.CategoryId=updateProduct.CategoryId
	}
	if sku!=""{
		productDetails.SKU=sku
	}
	if price>=0{
		productDetails.Price=price
	}
	productDetails.Specification=specification
	productDetails.Inventory.Quantity=quantity
	db.Save(&productDetails)
	c.JSON(200,productDetails)
}