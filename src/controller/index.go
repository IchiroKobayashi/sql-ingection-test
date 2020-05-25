package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sql-ingection-test/src/model"
)

// IndexGET displays application index page
func IndexGET(c *gin.Context) {

	c.HTML(200, "index.html", gin.H{})
}

func LoginPOST(c *gin.Context) {

	id := c.PostForm("id")
	pass := c.PostForm("pass")

	bool := model.FindByID(id, pass)

	if bool {
		data := model.GetAll()
		fmt.Println(data)

		c.HTML(200, "login.html", gin.H{
			"data": data,
		})
	} else {
		error := "ログインできません"

		c.HTML(200, "index.html", gin.H{"error": error})
	}

}


func NameSearchPOST(c *gin.Context) {

	name := c.PostForm("name")
	model.FindUser(name)
	c.HTML(200, "login.html", gin.H{
		"name": name,
	})
}

func NameCreatePOST(c *gin.Context) {

	name := c.PostForm("name")
	message := model.Create(name)
	if message == "作成完了" {
		c.HTML(200, "login.html", gin.H{
			"name": name,
		})
	} else {
		c.HTML(200, "login.html", gin.H{
			"message": message,
		})
	}
}