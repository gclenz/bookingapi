package room

import (
	"context"

	"github.com/gclenz/tinybookingapi/internal/app/domain/user"
)

type CreateRoom struct {
	repository     Repository
	userRepository user.Repository
}

func NewCreateRoom(repo Repository, ur user.Repository) *CreateRoom {
	return &CreateRoom{
		repository:     repo,
		userRepository: ur,
	}
}

func (cr *CreateRoom) Execute(
	staffID string,
	name string,
	singleBedCount int,
	doubleBedCount int,
	guestsLimit int,
	petFriendly bool,
	ctx context.Context,
) (*Room, error) {
	u, err := cr.userRepository.FindByID(staffID, ctx)
	if err != nil {
		return nil, err
	}

	if u.Role != user.StaffRole {
		return nil, ErrStaffOnlyCreateRoom
	}

	room := NewRoom(name, singleBedCount, doubleBedCount, guestsLimit, petFriendly)
	err = cr.repository.Create(room, ctx)
	if err != nil {
		return nil, ErrCreateRoom
	}
	return room, nil
}
