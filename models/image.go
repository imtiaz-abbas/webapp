package models

import (
	"github.com/satori/go.uuid"
)

// Image model
type Image struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key"`
	ImageURL  string    `json:"image_url"`
	ProductID uuid.UUID `json:"product_id" gorm:"foreign_key"`
}
