package models

import (
	"time"
)

// Event Information
type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"createdAt"`
	Images      []string  `json:"images"`
	Videos      []string  `json:"videos"`
	Reviews     []Review  `json:"reviews"`
	Cost        int       `json:"cost"`
	Type        string    `json:"type"`
}
