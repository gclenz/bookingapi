package booking

import "context"

type Repository interface {
	Create(booking *Booking, ctx context.Context) error
	CheckForOverlappingBooking(booking *Booking, ctx context.Context) ([]*Booking, error)
	FindBookingsByRoomID(roomID string, ctx context.Context) ([]*Booking, error) // The fromNow parameter allows reusability to staff members checking old bookings and for the frontend that will query for bookings starting from the current day and never in the past
	FindBookingsByCustomerID(customerID string, ctx context.Context) ([]*Booking, error)
}
