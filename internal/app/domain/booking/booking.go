package booking

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID         string    `json:"id"`
	RoomID     string    `json:"roomId"`
	CustomerID string    `json:"customerId"`
	StartOn    time.Time `json:"startOn"`
	EndOn      time.Time `json:"endOn"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func NewBooking(customerID string, roomID string, startOn time.Time, endOn time.Time) *Booking {
	now := time.Now()
	return &Booking{
		ID:         uuid.NewString(),
		RoomID:     roomID,
		CustomerID: customerID,
		StartOn:    startOn,
		EndOn:      endOn,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

func BuildBooking(id string, customerID string, roomID string, startOn time.Time, endOn time.Time, createdAt time.Time, updatedAt time.Time) *Booking {
	return &Booking{
		ID:         id,
		RoomID:     roomID,
		CustomerID: customerID,
		StartOn:    startOn,
		EndOn:      endOn,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}
}
