package bookingdb

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/gclenz/tinybookingapi/internal/app/domain/booking"
)

type BookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) booking.Repository {
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
		&booking.StartOn,
		&booking.EndOn,
		&booking.CreatedAt,
		&booking.UpdatedAt,
	)

	if err != nil {
		slog.Error("BookingRepository(Create)", "error", err)
		return err
	}

	return nil
}

func (br *BookingRepository) CheckForOverlappingBooking(bk *booking.Booking, ctx context.Context) ([]*booking.Booking, error) {
	rows, err := br.db.QueryContext(ctx, SelectOverlappingBookingQuery, &bk.RoomID, &bk.StartOn, &bk.EndOn)
	if err != nil {
		slog.Error("BookingRepository(CheckForOverlappingBooking)", "error", err)
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	if rows.Err() != nil {
		slog.Error("BookingRepository(CheckForOverlappingBooking)", "error", err)
		return nil, err
	}

	bookings := make([]*booking.Booking, 0)

	for rows.Next() {
		var bkg booking.Booking
		err := rows.Scan(
			&bkg.ID,
			&bkg.RoomID,
			&bkg.CustomerID,
			&bkg.StartOn,
			&bkg.EndOn,
			&bkg.CreatedAt,
			&bkg.UpdatedAt,
		)
		if err != nil {
			slog.Error("BookingRepository(CheckForOverlappingBooking) error:", err)
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
		slog.Error("BookingRepository(FindBookingsByRoomID)", "error", err)
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	if rows.Err() != nil {
		slog.Error("BookingRepository(FindBookingsByRoomID)", "error", err)
		return nil, err
	}

	bookings := make([]*booking.Booking, 0)

	for rows.Next() {
		var bkg booking.Booking
		err := rows.Scan(
			&bkg.ID,
			&bkg.RoomID,
			&bkg.CustomerID,
			&bkg.StartOn,
			&bkg.EndOn,
			&bkg.CreatedAt,
			&bkg.UpdatedAt,
		)
		if err != nil {
			slog.Error("BookingRepository(FindBookingsByRoomID)", "error", err)
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
		slog.Error("BookingRepository(FindBookingsByCustomerID)", "error", err)
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	if rows.Err() != nil {
		slog.Error("BookingRepository(FindBookingsByCustomerID)", "error", err)
		return nil, err
	}

	bookings := make([]*booking.Booking, 0)

	for rows.Next() {
		var bkg booking.Booking
		err := rows.Scan(
			&bkg.ID,
			&bkg.RoomID,
			&bkg.CustomerID,
			&bkg.StartOn,
			&bkg.EndOn,
			&bkg.CreatedAt,
			&bkg.UpdatedAt,
		)
		if err != nil {
			slog.Error("BookingRepository(FindBookingsByCustomerID)", "error", err)
			return nil, err
		}

		bookings = append(bookings, &bkg)
	}

	return bookings, nil
}
