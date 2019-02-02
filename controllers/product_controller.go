package controllers

import (
	"fmt"

	"github.com/webapp/concerns"
	"github.com/webapp/utilities"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/webapp/db"
	"github.com/webapp/models"
)

// GetAllProducts method
func GetAllProducts(c *gin.Context) {
	c.JSON(200, gin.H{"message": "got request"})
}

// AddProductsRequestBody type
type AddProductsRequestBody struct {
	Products []ProductRequestBody `json:"products"`
}

// ProductRequestBody TYPE
type ProductRequestBody struct {
	concerns.Storable
	Name        string            `json:"name"`
	Price       int               `json:"price"`
	Description string            `json:"description"`
	CategoryID  uuid.UUID         `json:"category_id"`
	Quantity    int               `json:"quantity"`
	Status      string            `json:"status"`
	Locations   []models.Location `json:"locations"`
}

// AddProduct method
func AddProduct(c *gin.Context) {
	request := ProductRequestBody{}
	if err := c.Bind(&request); err != nil {
		fmt.Println(" ==== UNABLE TO BIND REQUEST ====")
		c.JSON(400, gin.H{"message": "ERROR"})
	}
	request.ID = utilities.GenerateID()
	if err := db.Get().Table("products").Create(&request).Error; err != nil {
		fmt.Println(" ==== UNABLE TO BIND REQUEST ====")
		c.JSON(400, err)
	}

	c.JSON(200, request)
}
