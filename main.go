package main

import (
	"log"
	"os"
	"github.com/webapp/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.Use(gin.Logger())
	router.GET("/products", controllers.GetAllProducts)
	router.GET("/products/:id", controllers.GetProduct)

	router.Run(":" + port)
}