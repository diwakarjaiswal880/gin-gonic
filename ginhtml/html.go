//basic HTML service
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("template/*.html")
	r.GET("/", displayString)
	r.GET("/user/:name", userInfo)
	r.GET("/html", displayHtml)
	r.Run(":8080")
}

type information struct {
	Name string
	Age  int
}

func displayHtml(c *gin.Context) {
	info := information{Name: "Diwakar Jaiswal", Age: 24}
	c.HTML(http.StatusOK, "home.html", info)
}
func userInfo(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"User Name is": name,
	})
}
func displayString(c *gin.Context) {
	c.String(http.StatusOK, "HELLO WORLD")
}
