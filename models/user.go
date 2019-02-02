package models

import (
	"time"

	"github.com/satori/go.uuid"
)

// User type
type User struct {
	ID             uuid.UUID  `json:"id" gorm:"primary_key"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `sql:"index"`
	Name           string     `json:"name"`
	Gender         string     `json:"gender"`
	Age            int        `json:"age"`
	Address        string     `json:"address"`
	PhoneNumber    string     `json:"phone_number"`
	AltPhoneNumber string     `json:"alt_phone_number"`
	EmailID        string     `json:"email_id"`
	Image          Attachment `json:"image"`
	AadhaarID      string     `json:"aadhaar_id"`
	PanNumber      string     `json:"pan_number"`
	DrivingLicense string     `json:"driving_license"`
	GeneratedID    string     `json:"generated_id"`
	Orders         []Order    `json:"orders"`
}
