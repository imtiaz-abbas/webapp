package controllers

import (
	"fmt"

	"github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
	"github.com/webapp/db"
	"github.com/webapp/utilities"
)

// CategoryCreateRequest type
type CategoryCreateRequest struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// AddCategory method
func AddCategory(c *gin.Context) {
	request := CategoryCreateRequest{}
	if err := c.Bind(&request); err != nil {
		fmt.Println(" ==== UNABLE TO BIND REQUEST ====")
		c.JSON(400, gin.H{"message": "ERROR"})
	}
	request.ID = utilities.GenerateID()
	if err := db.Get().Table("categories").Create(&request).Error; err != nil {
		fmt.Println(" ==== UNABLE TO BIND REQUEST ====")
		c.JSON(400, err)
	}

	c.JSON(200, request)
}
