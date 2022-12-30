package test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Akash-Parmar-0917/ecommerce/pkg/controllers"
	"github.com/Akash-Parmar-0917/ecommerce/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


func TestAddToCart(t *testing.T){
	//test1
	w:= httptest.NewRecorder()
	c := GetTestGinContext(w)
	params := []gin.Param{
		{
			Key:   "cartID",
			Value: "/1",
		},
	}
	c.Params = params
	data:=`{
		"id":1,
		"quantity":1
	  }`
	MockJsonPost(c,data)
	controllers.AddToCart(c)
	assert.EqualValues(t, http.StatusNotFound, w.Code)
	assert.Equal(t,"cart with cart_id 1 not found",w.Body.String())
	
	//test2
	w= httptest.NewRecorder()
	c = GetTestGinContext(w)
	params = []gin.Param{
		{
			Key:   "cartID",
			Value: "/",
		},
	}
	c.Params = params
	data=`{
		"id":100,
		"quantity":1
	  }`
	MockJsonPost(c,data)
	controllers.AddToCart(c)
	assert.EqualValues(t, http.StatusBadRequest, w.Code)
	assert.Equal(t,`{"error":"please check details for below proudct_ids","product_id":[100]}`,w.Body.String())
	
	//test3
	
	w= httptest.NewRecorder()
	c = GetTestGinContext(w)
	params = []gin.Param{
		{
			Key:   "cartID",
			Value: "/",
		},
	}
	c.Params = params
	data=`{
		"id":1,
		"quantity":1
	  }`
	MockJsonPost(c,data)
	controllers.AddToCart(c)
	assert.EqualValues(t, http.StatusOK, w.Code)
	value, err := ioutil.ReadAll(w.Body)
	assert.NoError(t,nil,err)
	
	var cart models.Cart
	err=json.Unmarshal(value,&cart)
	assert.NoError(t,nil,err)
	if(cart.ID==0 ){
		t.Errorf("expectd category but got empty category %+v",cart)
	}
	if(len(cart.Product)==0){
		t.Error("expected cart with one product but got zero product")
	}
	if(cart.Price!=10){
		t.Errorf("expectd price 10 but got %v",cart.Price)
	}

	//test4
	w= httptest.NewRecorder()
	c = GetTestGinContext(w)
	params = []gin.Param{
		{
			Key:   "cartID",
			Value: "/1",
		},
	}
	c.Params = params
	data_newL:=struct{Id []int64 `json:"id"`;Quantity []int64 `json:"quantity"`}{
		Id: []int64{1}, 
		Quantity: []int64{1},
	}
	// var bt []byte
	bt,err:=json.Marshal(data_newL)
	assert.NoError(t,nil,err)
	assert.Equal(t,`{"id":[1],"quantity":[1]}`,string(bt))
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bt))
	// MockJsonPost(c,data)
	controllers.AddToCart(c)
	assert.EqualValues(t, http.StatusOK, w.Code)
	value, err = ioutil.ReadAll(w.Body)
	assert.NoError(t,nil,err)
	
	var cart2 models.Cart
	err=json.Unmarshal(value,&cart2)
	assert.NoError(t,nil,err)
	if(cart2.ID==0 ){
		t.Errorf("expectd category but got empty category %+v",cart2)
	}
	if(len(cart2.Product)==0){
		t.Error("expected cart with one product but got zero product")
	}
	if(cart2.Price!=20){
		t.Errorf("expectd price 30 but got %v",cart2.Price)
	}
}



// func TestGetCart(t *testing.T){

// }
