package booking

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customerId"`
	RoomID     string    `json:"roomId"`
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func NewBooking(customerID string, roomID string, start time.Time, end time.Time) *Booking {
	now := time.Now()
	return &Booking{
		ID:         uuid.NewString(),
		CustomerID: customerID,
		RoomID:     roomID,
		Start:      start,
		End:        end,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

func BuildBooking(id string, customerID string, roomID string, start time.Time, end time.Time, createdAt time.Time, updatedAt time.Time) *Booking {
	return &Booking{
		ID:         id,
		CustomerID: customerID,
		RoomID:     roomID,
		Start:      start,
		End:        end,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}
}
