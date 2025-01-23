package payment

import (
	"context"
)

type Gateway interface {
	Create(payment *Payment, ctx context.Context) error
	FindByID(externalID string, ctx context.Context) (*Payment, error)
}
