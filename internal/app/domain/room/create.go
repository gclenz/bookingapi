package room

import (
	"github.com/gclenz/tinybookingapi/internal/app/domain/user"
)

type CreateRoom struct {
	repository     Repository
	userRepository user.Repository
}

func (cr *CreateRoom) Execute(
	staffID string,
	name string,
	singleBedCount int,
	doubleBedCount int,
	guestsLimit int,
	arePetsAllowed bool,
) (*Room, error) {
	u, err := cr.userRepository.FindByID(staffID)
	if err != nil {
		return nil, err
	}

	if u.Role != user.StaffRole {
		return nil, ErrStaffOnlyCreateRoom
	}

	room := NewRoom(name, singleBedCount, doubleBedCount, guestsLimit, arePetsAllowed)
	err = cr.repository.Create(room)
	if err != nil {
		return nil, ErrCreateRoom
	}
	return room, nil
}
