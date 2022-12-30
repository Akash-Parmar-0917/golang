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


func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
	Header: make(http.Header),
	URL:    &url.URL{},
	}

	return ctx
}


func TestAddCategory(t *testing.T){
	w := httptest.NewRecorder()
	c := GetTestGinContext(w)
	data:=`{"category_name":"c1"}`
	MockJsonPost(c,data)
	controllers.AddCategory(c)

	assert.EqualValues(t, http.StatusOK, w.Code)
	value, err := ioutil.ReadAll(w.Body)
	assert.NoError(t,nil,err)
	
	var category models.Category
	err=json.Unmarshal(value,&category)
	assert.NoError(t,nil,err)
	if(category.ID==0 || category.CategoryName!="c1"){
		t.Errorf("expectd category but got empty category %+v",category)
	}
	// t.Errorf("got %+v",category)
}

func TestGetCategoryById(t *testing.T){
	w := httptest.NewRecorder()
	c := GetTestGinContext(w)
	params := []gin.Param{
		{
			Key:   "categoryID",
			Value: "1",
		},
	}
	u := url.Values{}
	MockJsonGet(c, params, u)
	controllers.GetCategoryById(c)
	assert.EqualValues(t, http.StatusOK, w.Code)


	value, err := ioutil.ReadAll(w.Body)
	assert.NoError(t,nil,err)
	
	var category models.Category
	err=json.Unmarshal(value,&category)
	assert.NoError(t,nil,err)

	if(category.ID!=1){
		t.Errorf("expected category_id 1 got %v",category.ID)
	}

	w = httptest.NewRecorder()
	c = GetTestGinContext(w)
	params = []gin.Param{
		{
			Key:   "categoryID",
			Value: "200000000000",
		},
	}
	MockJsonGet(c, params, u)
	controllers.GetCategoryById(c)
	assert.EqualValues(t, http.StatusNotFound, w.Code)

	value, err = ioutil.ReadAll(w.Body)
	assert.NoError(t,nil,err)
	var category2 models.Category
	err=json.Unmarshal(value,&category2)
	assert.NoError(t,nil,err)

	if(category2.ID!=0){
		t.Errorf("expected empty category struct but got %+v",category2)
	}
}
