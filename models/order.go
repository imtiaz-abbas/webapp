package models

import (
	"time"

	"github.com/satori/go.uuid"
	"github.com/webapp/concerns"
)

// Order type
type Order struct {
	concerns.Storable
	ProductIDS      []uuid.UUID `json:"product_ids"`
	UserID          uuid.UUID   `json:"user_id"`
	BookingDuration int         `json:"booking_duration"`
	RentAmount      int         `json:"rent_amount"`
	ShippingCharges int         `json:"shipping_charges"`
	Deposit         int         `json:"deposit"`
	PaymentMode     string      `json:"payment_mode"`
	StartDate       time.Time   `json:"start_date"`
	EndDate         time.Time   `json:"end_date"`
	Status          string      `json:"status"`
	DeliveryAddress string      `json:"delivery_address"`
}
