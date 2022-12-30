package routes

import (
	"github.com/Akash-Parmar-0917/ecommerce/pkg/controllers"
	"github.com/gin-gonic/gin"
)



var RegisterEcommerceRoutes = func(router *gin.Engine){
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
}