package main

import (
	"html/template"

	"github.com/Akash-Parmar-0917/ecommerce/pkg/config"
	"github.com/Akash-Parmar-0917/ecommerce/pkg/models"
	"github.com/Akash-Parmar-0917/ecommerce/pkg/routes"
	"github.com/Akash-Parmar-0917/ecommerce/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// func loadHTMLGlob(engine *gin.Engine, pattern string) {
//     funcMap := template.FuncMap{
// 		"jsonToMap":utils.JsonToMap, // A custom template function
//     }
//     //if gin.IsDebugging() {
//     //  debugPrintLoadTemplate(template.Must(template.New("").Funcs(funcMap).ParseGlob(pattern)))
//     //  engine.HTMLRender = render.HTMLDebug{Glob: pattern}
//     //} else {
//     templ := template.Must(template.New("").Funcs(funcMap).ParseGlob(pattern))
//     engine.SetHTMLTemplate(templ)
//     //}
// }



func Migrate(db *gorm.DB){
	db.Migrator().DropTable(&models.Cart{},&models.CartProduct{})
	db.SetupJoinTable(&models.Cart{}, "Product", &models.CartProduct{})
	db.AutoMigrate(&models.Product{},&models.Inventory{},&models.Category{},&models.Cart{},&models.CartProduct{})
}

func main(){
	db:=config.Connect();
	Migrate(db)
	r:= gin.Default();
	r.UseRawPath = true
	r.UnescapePathValues = false
	// loadHTMLGlob(r,"../../templates/*")
	r.SetFuncMap(template.FuncMap{
		"jsonToMap":utils.JsonToMap,
	})
	r.Static("/static","./static")
	r.LoadHTMLGlob("templates/*")
	routes.RegisterEcommerceRoutes(r)
	r.Run(":8080")
	}