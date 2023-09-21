package booking

import "errors"

var (
	ErrOverlappingBooking = errors.New("there is already a booking between the chosen dates")
	ErrCreateBooking      = errors.New(
		"something went wrong while trying to create a new booking. Please, try again. If the error persists, contact the development team",
	)
	ErrBookingInThePast         = errors.New("you can't book a room in the past")
	ErrEndDateBeforeStartAtDate = errors.New("you can't have an end date before the start date")
)
