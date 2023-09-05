package bookingdb

import (
	"context"
	"database/sql"

	"github.com/gclenz/tinybookingapi/internal/app/domain/booking"
)

type BookingRepository struct {
	db *sql.DB
}

func (br *BookingRepository) NewBookingRepository(db *sql.DB) booking.Repository {
	return &BookingRepository{
		db: db,
	}
}

func (br *BookingRepository) Create(booking *booking.Booking, ctx context.Context) error {
	_, err := br.db.ExecContext(
		ctx,
		InsertBookingQuery,
		&booking.ID,
		&booking.RoomID,
		&booking.CustomerID,
		&booking.Start,
		&booking.End,
		&booking.CreatedAt,
		&booking.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (br *BookingRepository) CheckForOverlappingBooking(bk *booking.Booking, ctx context.Context) ([]*booking.Booking, error) {
	rows, err := br.db.QueryContext(ctx, SelectOverlappingBookingQuery, &bk.RoomID, &bk.Start, &bk.End)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	if rows.Err() != nil {
		return nil, err
	}

	bookings := make([]*booking.Booking, 0)

	for rows.Next() {
		var bkg booking.Booking
		err := rows.Scan(&bkg)
		if err != nil {
			return nil, err
		}

		bookings = append(bookings, &bkg)
	}

	return bookings, nil
}

func (br *BookingRepository) FindBookingsByRoomID(roomID string, ctx context.Context) ([]*booking.Booking, error) {
	rows, err := br.db.QueryContext(
		ctx,
		SelectBookingsByRoomIDQuery,
		roomID,
	)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	if rows.Err() != nil {
		return nil, err
	}

	bookings := make([]*booking.Booking, 0)

	for rows.Next() {
		var bkg booking.Booking
		err := rows.Scan(&bkg)
		if err != nil {
			return nil, err
		}

		bookings = append(bookings, &bkg)
	}

	return bookings, nil
}

func (br *BookingRepository) FindBookingsByCustomerID(customerID string, ctx context.Context) ([]*booking.Booking, error) {
	rows, err := br.db.QueryContext(
		ctx,
		SelectBookingsByCustomerIDQuery,
		customerID,
	)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	if rows.Err() != nil {
		return nil, err
	}

	bookings := make([]*booking.Booking, 0)

	for rows.Next() {
		var bkg booking.Booking
		err := rows.Scan(&bkg)
		if err != nil {
			return nil, err
		}

		bookings = append(bookings, &bkg)
	}

	return bookings, nil
}
