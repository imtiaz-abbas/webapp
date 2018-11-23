package models

import (
	"time"
)

type Blog struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Article   string    `json:"article"`
	Images    []string  `json:"images"`
	Videos    []string  `json:"videos"`
	CreatedAt time.Time `json:"createdAt"`
}
