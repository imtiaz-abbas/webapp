package controllers

import (
	"github.com/webapp/models"
	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	allProducts := models.GetAllProducts()
	c.JSON(200, allProducts)
}
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, models.GetProduct(id))
}