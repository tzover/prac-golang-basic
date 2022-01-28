package main

import (
	"fmt"
	"gin/calc"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello golang from docker!")

	calc.Sum()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello golang from docker!",
		})
	})

	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")

		fileFind()

		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.POST("/user/:name/*action", func(c *gin.Context) {
		path := c.FullPath()
		c.String(http.StatusOK, path)
	})
	router.Run(":8080")
}

func fileFind() {
	var fileName string
	files, _ := ioutil.ReadDir("./public/text")
	for _, file := range files {
		fileName = file.Name()
		fmt.Println(fileName)
	}
}
