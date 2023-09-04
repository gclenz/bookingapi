package booking

import (
	"time"

	"github.com/gclenz/tinybookingapi/internal/app/domain/room"
	"github.com/gclenz/tinybookingapi/internal/app/domain/user"
)

type CreateBooking struct {
	repository     Repository
	roomRepository room.Repository
	userRepository user.Repository
}

func (cb *CreateBooking) Execute(customerID string, roomID string, start time.Time, end time.Time) (*Booking, error) {
	_, err := cb.roomRepository.FindByID(roomID)
	if err != nil {
		return nil, room.ErrRoomNotFound
	}

	_, err = cb.userRepository.FindByID(customerID)
	if err != nil {
		return nil, user.ErrUserNotFound
	}

	booking := NewBooking(customerID, roomID, start, end)
	overlappingBookings, err := cb.repository.CheckForOverlappingBooking(booking)
	if err != nil {
		return nil, err
	}

	if len(overlappingBookings) > 0 {
		return nil, ErrOverlappingBooking
	}

	err = cb.repository.Create(booking)
	if err != nil {
		return nil, err
	}

	return booking, nil
}
