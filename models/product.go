package models

import "time"

type Product struct {
	ID               string      `json:"id"`
	Name             string      `json:"name"`
	Description      string      `json:"description"`
	Category         Category    `json:"category"`
	Images           []string    `json:"images"`
	Value            int         `json:"value"`
	Weight           int         `json:"weight"`
	UnavailableDates []time.Time `json:"unavailableDates"`
	TotalRating      int         `json:"totalRating"`
	Reviews          []Review    `json:"reviews"`
}
