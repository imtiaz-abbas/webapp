package models

import (
	"time"
)

type Review struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Images    []string  `json:"images"`
	Review    string    `json:"review"`
	CreatedAt time.Time `json:"createdAt"`
	Rating    int       `json:"rating"`
	UserID    string    `json:"userId"`
}
