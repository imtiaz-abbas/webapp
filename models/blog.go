package models

import (
	"time"
)

// Blog holds the basic information about blog
type Blog struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Article   string    `json:"article"`
	Images    []string  `json:"images"`
	Videos    []string  `json:"videos"`
	CreatedAt time.Time `json:"createdAt"`
}
