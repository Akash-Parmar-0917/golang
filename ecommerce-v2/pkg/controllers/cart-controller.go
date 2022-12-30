package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Akash-Parmar-0917/ecommerce/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)



func GetCart(c *gin.Context){
	temp:=c.Param("cartID")
	cartID,err:=strconv.ParseInt(temp,0,0)
	if err!=nil{
		fmt.Println("error while Parsing in GetCart", err)
		return
	}
	cartdetail,_:=models.GetCartById(cartID)
	if(cartdetail.ID==0){
		c.String(http.StatusBadRequest,"cart with cart_id %v not found",cartID)
		return
	}
	c.JSON(http.StatusOK,cartdetail)
}

func AddToCart(c *gin.Context){
	temp:=c.Param("cartID")
	var cartID int64 =0
	if(temp!="/"){
		var err error
		cartID,_=strconv.ParseInt(temp[1:],0,0)
		if err!=nil{
			fmt.Println("error while Parsing in GetCart", err)
			return
		}
	}
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("error while reading data for addToCart")
		return
	}
	m:=make(map[string]interface{})
	p:=make(map[int64]int64)
	var quantity []int64
	json.Unmarshal(jsonData,&m)

	switch y := m["quantity"].(type) {
    case []interface{}:
	    for _, i := range y {
            quantity=append(quantity,int64(i.(float64)))
        }
    case float64:
        quantity=append(quantity, int64(y))
    default:
        fmt.Printf("Unsupported type of quantity: %T\n", y)
		return
    }
	switch x := m["id"].(type) {
    case []interface{}:
	    for i, val := range x {
            p[int64(val.(float64))]=quantity[i]
        }
    case float64:
        p[int64(x)]=quantity[0]
    default:
        fmt.Printf("Unsupported type of id: %T\n", x)
		return
    }
	
	var price int64=0
	var badProductId []int64
	for id,quantity:=range p{
		product,err:=models.CheckProductForCart(id,quantity,cartID)
		if(err==nil){
			price+=product.Price*p[int64(product.ID)]
		}else{
			badProductId=append(badProductId,id)	
		}
	}
	if(len(badProductId)!=0 ){
		fmt.Println("line 87 in cart-controller:",badProductId)
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"please check details for below proudct_ids",
			"product_id":badProductId,
		})
		return
	}
	var cart *models.Cart
	var db *gorm.DB
	if(cartID==0){
		cart,db=models.CreateCart(&models.Cart{})
		cartID=int64(cart.ID)
	}else{
		cart,db=models.GetCartById(cartID)
		if(cart.ID==0){
			c.String(http.StatusNotFound,"cart with cart_id %v not found",cartID)
			return
		}
	}
	for id,quantity:=range p{
		models.AddProductToCart(cartID,int64(id),quantity)	
	}
	db.Model(&models.Cart{}).Where("id=?",cartID).Update("price",cart.Price+price)
	cart,_=models.GetCartById(cartID)
	c.JSON(http.StatusOK,*cart)
}

func UpdateCart(c *gin.Context){
	temp:=c.Param("cartID")
	cartID,err:=strconv.ParseInt(temp,0,0)
	if err!=nil{
		fmt.Println("error while Parsing in GetCart", err)
		return
	}

	var cart *models.Cart
	cart,_=models.GetCartById(cartID)
	if(cart.ID==0){
		c.String(http.StatusBadRequest,"Invalid cart_id %v",cartID)
		return
	}

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("error while reading data for UpdateCart")
		return
	}

	m:=make(map[string]interface{})
	p:=make(map[int64]int64)
	var quantity []int64
	json.Unmarshal(jsonData,&m)

	switch y := m["quantity"].(type) {
    case []interface{}:
	    for _, i := range y {
            quantity=append(quantity,int64(i.(float64)))
        }
    case float64:
        quantity=append(quantity, int64(y))
    default:
        fmt.Printf("Unsupported type of quantity: %T\n", y)
		return
    }
	switch x := m["product_id"].(type) {
    case []interface{}:
	    for i, val := range x {
            p[int64(val.(float64))]=quantity[i]
        }
    case float64:
        p[int64(x)]=quantity[0]
    default:
        fmt.Printf("Unsupported type of id: %T\n", x)
		return
    }

	var badProductId []int64
	for id,quantity:=range p{
		_,err:=models.CheckProductForCart(id,quantity,0)
		if(err!=nil){
			fmt.Println(err)
			badProductId=append(badProductId,id)	
		}
	}
	if(len(badProductId)!=0 ){
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"please check details for below proudct_ids",
			"product_id":badProductId,
		})
		return
	}

	for id,quantity:=range p{
		cart,err=models.UpdateCart(cartID,id,quantity)
		fmt.Println(err)
	}
	c.JSON(http.StatusOK,*cart)
}

func DeleteFromCart(c *gin.Context){
	temp:=c.Param("cartID")
	cartID,_:=strconv.ParseInt(temp,0,0)
	cart,_:=models.GetCartById(cartID)
	if(cart.ID==0){
		c.String(http.StatusBadRequest,"Invalid cart_id %v",cartID)
		return
	}
	temp=c.Param("productID")
	var productID int64 =0
	if(temp!="/"){
		var err error
		productID,_=strconv.ParseInt(temp[1:],0,0)
		if err!=nil{
			fmt.Println("error while Parsing in GetCart", err)
			return
		}
	}
	if(productID==0){
		models.DeleteCart(cartID)
		c.String(http.StatusOK,"removed cart_id %v",productID,cartID)
	}else{
		models.DeleteCartProduct(cartID,productID)
		c.String(http.StatusOK,"item with product_id %v removed from cart_id %v",productID,cartID)
	}

	
}