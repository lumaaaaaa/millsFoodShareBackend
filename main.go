package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
// TODO: Get this from the db in the future
)

func initialize() {
	// TODO: Read from database and create connection
	fmt.Println("Reading values from database...")
}

func main() {
	initialize()

	r := gin.Default()
	r.GET("/api/items", func(c *gin.Context) {
		// TODO: read needed items from database
		neededItems := []string{"Canned Food", "Kitty Litter", "Toothpaste"}
		c.JSON(http.StatusOK, gin.H{
			"neededItems": neededItems,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
