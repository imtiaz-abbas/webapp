package models

import (
	"github.com/satori/go.uuid"
)

// Product model
type Product struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key"`
	Name        string    `json:"name" gorm:"unique"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	Images      []Image   `json:"images"`
}
