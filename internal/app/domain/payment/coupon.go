package payment

import (
	"time"

	"github.com/google/uuid"
)

type Coupon struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	Mode      string    `json:"mode"`
	Value     float64   `json:"grossAmount"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewCoupon(code, mode string, value float64, expiresAt time.Time) *Coupon {
	now := time.Now()
	return &Coupon{
		ID:        uuid.NewString(),
		Code:      code,
		Mode:      mode,
		Value:     value,
		ExpiresAt: expiresAt,
		CreatedAt: now,
	}
}

func BuildCoupon(id, code, mode string, value float64, expiresAt, createdAt time.Time) *Coupon {
	return &Coupon{
		ID:        id,
		Code:      code,
		Mode:      mode,
		Value:     value,
		ExpiresAt: expiresAt,
		CreatedAt: createdAt,
	}
}
