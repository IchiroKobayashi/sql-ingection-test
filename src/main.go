package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sql-ingection-test/src/controller"
)

func main() {


	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", controller.IndexGET)

	router.POST("/login", controller.LoginPOST)



	router.POST("/search", controller.NameSearchPOST)
	router.POST("/new", controller.NameCreatePOST)
	router.Run(":8080")

}