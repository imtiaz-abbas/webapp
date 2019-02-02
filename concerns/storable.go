package concerns

import (
	"time"

	"github.com/satori/go.uuid"
)

// Storable Encapsulates mandatory fields in any table record.
type Storable struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
