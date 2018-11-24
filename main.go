package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/webapp/controllers"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		// log.Fatal("$PORT must be set")
		port = "8080"
	}

	router := gin.Default()
	router.Use(gin.Logger())
	router.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "OK"}) })
	router.GET("/products", controllers.GetAllProducts)
	router.GET("/products/:id", controllers.GetProduct)

	router.Run(":" + port)
}
