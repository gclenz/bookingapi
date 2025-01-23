package payment

import (
	"time"

	"github.com/google/uuid"
)

type Coupon struct {
	ID        string    `json:"id"`
	Category  string    `json:"category"`
	Value     float64   `json:"grossAmount"`
	ExpiresIn time.Time `json:"expiresIn"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewCoupon(category string, value float64, expiresIn time.Time) *Coupon {
	now := time.Now()
	return &Coupon{
		ID:        uuid.NewString(),
		Category:  category,
		Value:     value,
		ExpiresIn: expiresIn,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func BuildCoupon(id, category string, value float64, expiresIn, createdAt, updatedAt time.Time) *Coupon {
	return &Coupon{
		ID:        id,
		Category:  category,
		Value:     value,
		ExpiresIn: expiresIn,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
