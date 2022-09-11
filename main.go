package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//untuk iport pkg
	router := gin.Default()

	//membuat request
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Vivy Cahyani",
			"bio":  "Pelajar",
		})
	})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title":    "Hello World",
			"subtitle": "Beajar golang nich",
		})
	})

	//mengganti port route
	router.Run(":8888")
}
