package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//untuk iport pkg
	router := gin.Default()

	//membuat route
	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)

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

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	//query string
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}
