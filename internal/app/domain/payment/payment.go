package payment

import (
	"time"

	"github.com/google/uuid"
)

// id uuid primary key default(gen_random_uuid()) not null,
// external_id varchar(50) not null,
// booking_id uuid references bookings(id) not null,
// coupon_id uuid references coupons(id),
// gross_amount decimal(10,2) not null,
// net_amount decimal(10,2) not null,
// status varchar(50) not null,
// created_at timestamp default(now()) not null,
// updated_at timestamp default(now()) not null

type Payment struct {
	ID          string    `json:"id"`
	ExternalID  string    `json:"externalId"`
	BookingID   string    `json:"bookingId"`
	CouponID    string    `json:"couponId,omitempty"`
	GrossAmount float64   `json:"grossAmount"`
	NetAmount   float64   `json:"netAmount"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewPayment(externalID, bookingID, couponID string, grossAmount, netAmount float64) *Payment {
	now := time.Now()
	return &Payment{
		ID:          uuid.NewString(),
		ExternalID:  externalID,
		BookingID:   bookingID,
		CouponID:    couponID,
		GrossAmount: grossAmount,
		NetAmount:   netAmount,
		Status:      "pending",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func BuildPayment(id, externalID, bookingID, couponID, status string, grossAmount, netAmount float64, createdAt, updatedAt time.Time) *Payment {
	return &Payment{
		ID:          id,
		ExternalID:  externalID,
		BookingID:   bookingID,
		CouponID:    couponID,
		GrossAmount: grossAmount,
		NetAmount:   netAmount,
		Status:      status,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
