package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var (
	// TODO: Get this from the db in the future
	donations *Donations
)

func initialize() {
	// TODO: Read from database and create connection
	fmt.Println("Setting up database...")
	var err error
	donations, err = NewDonations()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected!")
}

func main() {
	initialize()

	r := gin.Default()
	r.GET("/api/items", func(c *gin.Context) {
		var err error
		var itemsResponse ItemsResponse
		i := 0
		for err == nil {
			var donation Donation
			donation, err = donations.Retrieve(i)
			if err != nil {
				log.Println("Reached end of index at row number: " + strconv.Itoa(i))
			}

			if donation.Name != "" {
				itemsResponse.Items = append(itemsResponse.Items, donation)
			}

			i++
		}

		if len(itemsResponse.Items) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status": 404,
			})
		} else {
			c.JSON(http.StatusOK, itemsResponse)
		}

	})

	r.POST("/api/items", func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "invalid data",
			})
			return
		}

		var donation Donation
		err = json.Unmarshal(jsonData, &donation)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "invalid data",
			})
			return
		}

		_, err = donations.Insert(donation)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status": "invalid data",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Donation": donation,
			"Status":   "added",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
