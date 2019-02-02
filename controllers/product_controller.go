package controllers

import (
	"github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
	"github.com/webapp/db"
	"github.com/webapp/models"
)

// GetAllProducts method
func GetAllProducts(c *gin.Context) {
	c.JSON(200, gin.H{"message": "got request"})
}

type AddProductsRequestBody struct {
	Products []ProductRequestBody `json:"products"`
}

type ProductRequestBody struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

// AddProducts method
func AddProducts(c *gin.Context) {
	request := &AddProductsRequestBody{}
	c.Bind(&request)

	products := make([]models.Product, 0)

	for _, reqProduct := range request.Products {
		product := models.Product{}
		product.Name = reqProduct.Name
		product.Price = reqProduct.Price
		id, err := uuid.NewV4()
		if err != nil {
			c.JSON(400, gin.H{"message": "something went wrong"})
			return
		}
		product.ID = id
		product.Description = reqProduct.Description
		products = append(products, product)
		if err := db.Get().Create(&product).Error; err != nil {
			c.JSON(400, gin.H{"message": "Invalid Request"})
			return
		}
	}

	c.JSON(200, products)
}
