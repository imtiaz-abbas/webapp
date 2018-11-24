package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/webapp/models"
)

func GetAllProducts(c *gin.Context) {
	err, allProducts := models.GetAllProducts()
	if err == 0 {
		c.JSON(404, gin.H{"error": "Error Finding Products"})
	} else {
		c.JSON(200, allProducts)
	}
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	product := models.Product{}
	statusCode := product.GetProduct(id)
	if statusCode == 1 {
		c.JSON(200, product)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Product with ID found"})
	}
}
