package models

import (
	"github.com/satori/go.uuid"
	"github.com/webapp/concerns"
)

// Product model
type Product struct {
	concerns.Storable
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Price       int          `json:"price"`
	Locations   []Location   `json:"locations"`
	Images      []Attachment `json:"images"`
	GeneratedID string       `json:"generated_id"`
	CategoryID  uuid.UUID    `json:"categories"`
	Status      string       `json:"status" gorm:"default:'available'"`
	Orders      []Order      `json:"orders"`
	Quantity    int          `json:"quantity"`
}

// Location type
type Location struct {
	ID      uuid.UUID `json:"id" gorm:"primary_key"`
	Name    string    `json:"name"`
	Pincode string    `json:"pincode"`
}

// Attachment type
type Attachment struct {
	ID             uuid.UUID `json:"id"`
	URL            string    `json:"attachment_url"`
	AttachmentType string    `json:"attachment_type"`
}

// Category type
type Category struct {
	ID       uuid.UUID   `json:"id"`
	Products []Product   `json:"products"`
	Name     string      `json:"name"`
	Image    *Attachment `json:"image"`
}
