package test

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"testing"
	"text/template"

	"github.com/Akash-Parmar-0917/ecommerce/pkg/config"
	"github.com/Akash-Parmar-0917/ecommerce/pkg/controllers"
	"github.com/Akash-Parmar-0917/ecommerce/pkg/models"
	"github.com/Akash-Parmar-0917/ecommerce/pkg/utils"
	unitTest "github.com/Valiben/gin_unit_test"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	// initialize the router
	router = gin.Default()

	// Handlers for testing
	router.POST("/product",controllers.CreateProduct)
	router.GET("/addproduct",controllers.AddProductPage)
	router.GET("/products/*page",controllers.GetProducts)
	router.GET("/updateproduct",controllers.UpdateProductPage)
	router.GET("/product/:productID",controllers.GetProductByID)
	router.POST("/product/:productID",controllers.UpdateProduct)
	router.DELETE("/product/:productID",controllers.DeleteProduct)


	router.POST("/addcategory",controllers.AddCategory)
	router.DELETE("/category/:categoryName",controllers.DeleteCategory)
	router.GET("/categories",controllers.GetCategories)
	router.GET("/category/:categoryID",controllers.GetCategoryById)

	router.GET("/getcart/:cartID",controllers.GetCart)
	router.POST("/addtocart/*cartID",controllers.AddToCart)
	router.POST("/updatecart/:cartID",controllers.UpdateCart)
	router.DELETE("/removeitemfromcart/:cartID/*productID",controllers.DeleteFromCart)

	router.UseRawPath = true
	router.UnescapePathValues = false
	// loadHTMLGlob(router,"../../templates/*")
	router.SetFuncMap(template.FuncMap{
		"jsonToMap":utils.JsonToMap,
	})
	router.Static("/static","/static")
	router.LoadHTMLGlob("../templates/*")

	// Setup the router
	unitTest.SetRouter(router)
	newLog := log.New(os.Stdout, "", log.Llongfile|log.Ldate|log.Ltime)
	unitTest.SetLog(newLog)
}

func MockJsonGet(c *gin.Context, params gin.Params, u url.Values) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")
	// c.Set("user_id", 1)

	// set path params
	c.Params = params

	// set query params
	c.Request.URL.RawQuery = u.Encode()
}

func MockJsonPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	// c.Set("user_id", 1)

	// jsonbytes, err := json.Marshal(content)
	// if err != nil {
	// 	panic(err)
	// }

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBufferString(content.(string)))
	// c.Request.Body= bytes.NewBufferString(bodyData)
}


// func getModels() []interface{} {
// 	return []interface{}{&models.Cart{},&models.Product{},&models.Category{},&models.Inventory{},&models.CartProduct{}}
// }

func createSchema() error {
	db:=config.GetDB()
	db.SetupJoinTable(&models.Cart{}, "Product", &models.CartProduct{})
	db.AutoMigrate(&models.Product{},&models.Inventory{},&models.Category{},&models.Cart{},&models.CartProduct{})
	product:= models.Product{
		Name: "p1",
		Category: models.Category{CategoryName: "c1"},
		SKU: "12a",
		Price: 10,
		Inventory: models.Inventory{
			Quantity: 1000,
		},
	}
	db.Debug().Model(&models.Product{}).Create(&product)
	fmt.Println(product)
	return nil
}

func deleteSchema() error{
	db:=config.GetDB()
	db.Migrator().DropTable(&models.Product{},&models.Inventory{},&models.Category{},&models.CartProduct{},&models.Cart{})
	return nil
}


func TestMain(m *testing.M) {
	// user and password will need to match running postgres instance
	config.TestDBConnect()
	db,_:=config.GetDB().DB()
	defer db.Close()

	// Check if DB is connected
	if err := db.Ping(); err != nil {
		fmt.Print(err.Error())
		panic(err)
	}

	createSchema()

	log.Println("Database setup for test")
	exitVal := m.Run()
	log.Println("Database dropped after test")

	deleteSchema()

	os.Exit(exitVal)
}