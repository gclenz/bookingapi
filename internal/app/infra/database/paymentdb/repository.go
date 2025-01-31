package paymentdb

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/gclenz/tinybookingapi/internal/app/domain/payment"
)

type PaymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) payment.Repository {
	return &PaymentRepository{
		db: db,
	}
}

// Create implements payment.Repository.
func (pr *PaymentRepository) Create(payment *payment.Payment, ctx context.Context) error {
	//id, booking_id, external_id, coupon_id, gross_amount, net_amount, status, created_at, updated_at
	_, err := pr.db.ExecContext(
		ctx,
		InsertPaymentQuery,
		&payment.ID,
		&payment.BookingID,
		&payment.ExternalID,
		&payment.CouponID,
		&payment.GrossAmount,
		&payment.NetAmount,
		&payment.Status,
		&payment.CreatedAt,
		&payment.UpdatedAt,
	)

	if err != nil {
		slog.Error("PaymentRepository(Create) error:", err)
		return err
	}

	return nil
}

// FindByID implements payment.Repository.
func (pr *PaymentRepository) FindByID(paymentID string, ctx context.Context) (*payment.Payment, error) {
	row := pr.db.QueryRowContext(
		ctx,
		SelectPaymentByIDQuery,
		paymentID,
	)
	err := row.Err()
	if err != nil {
		slog.Error("PaymentRepository(FindByID) error:", err)
		return nil, err
	}

	var payment payment.Payment

	err = row.Scan(
		&payment.ID,
		&payment.BookingID,
		&payment.ExternalID,
		&payment.CouponID,
		&payment.GrossAmount,
		&payment.NetAmount,
		&payment.Status,
		&payment.CreatedAt,
		&payment.UpdatedAt,
	)
	if err != nil {
		slog.Error("PaymentRepository(FindByID) error:", err)
		return nil, err
	}

	return &payment, nil
}
