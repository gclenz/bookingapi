package booking

type Repository interface {
	Create(booking *Booking) error
	CheckForOverlappingBooking(booking *Booking) ([]*Booking, error)
	FindBookingsByRoomID(roomID string) ([]*Booking, error) // The fromNow parameter allows reusability to staff members checking old bookings and for the frontend that will query for bookings starting from the current day and never in the past
	FindBookingsByCustomerID(customerID string) ([]*Booking, error)
}
