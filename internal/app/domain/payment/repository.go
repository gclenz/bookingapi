package payment

import (
	"context"
)

type Repository interface {
	Create(payment *Payment, ctx context.Context) error
	FindByID(paymentID string, ctx context.Context) (*Payment, error)
}
