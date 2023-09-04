package booking

import "errors"

var (
	ErrOverlappingBooking = errors.New("there is already a booking between the dates you choosed")
	ErrCreateBooking      = errors.New(
		"something went wrong while trying to create a new booking. Please, try again. If the error persists, contact the development team",
	)
)
