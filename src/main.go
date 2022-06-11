package main

import (
	"github.com/gin-gonic/gin"

	"dressme-scrapper/src/routes/flickr"

	"net/http"

	"dressme-scrapper/src/mongodb"
)

func main() {

	mongoClient := mongodb.Connect()

	flickr.SearchPhoto("4", "model", "Medium", "/home/olivier/dressme/images/flickr", mongoClient)

	router := gin.Default()

	router.POST("/search/flickr", func(c *gin.Context) {
		license := c.Query("license") // licenseId: "4, 5, 7, 9, 10"
		tag := c.Query("tag")
		quality := c.Query("quality") // "Low, Medium, High"
		path := c.Query("path")

		ids, err := flickr.SearchPhoto(license, tag, quality, path, mongoClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, ids)
	})

	// TODO: unwated tags POST
	// TODO: wanted tags POST

	router.Run("localhost:8080")
}
