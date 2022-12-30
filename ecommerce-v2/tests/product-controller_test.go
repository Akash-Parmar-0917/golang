package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Akash-Parmar-0917/ecommerce/pkg/controllers"
	"github.com/Akash-Parmar-0917/ecommerce/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// // func SetUpRouter() *gin.Engine{
// //     router := gin.Default()
// //     return router
// // }

func TestGetProductByID(t *testing.T){
	
	w:= httptest.NewRecorder()
	c := GetTestGinContext(w)
	params := []gin.Param{
		{
			Key:   "productID",
			Value: "1000000",
		},
	}
	u := url.Values{}
	MockJsonGet(c, params, u)
	controllers.GetProductByID(c)
	assert.EqualValues(t, http.StatusNotFound, w.Code)

	assert.Equal(t,`{"value":"record not found"}`,w.Body.String())
	// value, err = ioutil.ReadAll(w.Body)
	// assert.NoError(t,nil,err)
	
	// var product2 models.Product
	// err=json.Unmarshal(value,&product2)
	// assert.NoError(t,nil,err)

	// if(product2.ID!=0){
	// 	t.Errorf("expected empty product struct but got %+v",product2)
	// }
}


func TestCreateProduct(t *testing.T){
	w := httptest.NewRecorder()
	c := GetTestGinContext(w)
	
	v := url.Values{}
	v.Set("prodcutName", "p2")
	v.Add("key[]", "k1")
	v.Add("value[]", "v1")
	v.Add("key[]", "k2")
	v.Add("value[]", "v2")
	v.Set("sku", "123")
	v.Set("quantity", "10")
	v.Set("price", "10")
	v.Set("category", "1")
	c.Request.Method = "POST"
	c.Request.PostForm=v
	c.Request.Header.Set("Content-Type", "x-www-form-urlencoded")
	
	controllers.CreateProduct(c)

	value, err := ioutil.ReadAll(w.Body)
	assert.NoError(t,nil,err)
	
	var product models.Product
	err=json.Unmarshal(value,&product)
	assert.NoError(t,nil,err)
	if(product.ID==0){
		t.Errorf("expectd product but got empty product %+v",product)
	}
	if(product.Price!=10){
		t.Errorf("expectd product price 10 but got %v",product.Price)
	}
	if(product.SKU!="123"){
		t.Errorf("expectd product sku 123 but got %s",product.SKU)
	}
	if(product.Inventory.ID==0 || product.Inventory.Quantity!=10){
		t.Errorf("expectd product quantity 10 but got %v",product.Inventory.Quantity)
	}
	if(product.Category.ID==0 ){
		t.Errorf("expectd category with id non zero but got %+v",product.Category)
	}
	
	// t.Errorf("%+v",product)

}


// // var requestTests = []struct {
// // 	Url               string
// // 	Method            string
// // 	Headers           map[string]interface{}
// // 	Body              interface{}
// // 	ResponseCode      int
// // 	ExpectedData      interface{}
// // }{
// // 	"/product",
// // 	"POST",
// // 	nil,

// // }

// func (t *SuiteTest) TestCreateProduct(){

// }