package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//untuk iport pkg
	router := gin.Default()

	//membuat request
	router.GET("/", rootHandler)

	router.GET("/hello", helloHandler)

	//mengganti port route
	router.Run(":8888")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Vivy Cahyani",
		"bio":  "Pelajar",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":    "Hello World",
		"subtitle": "Beajar golang nich",
	})
}
