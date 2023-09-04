package room

import "errors"

var (
	ErrRoomNotFound        = errors.New("room not found")
	ErrStaffOnlyCreateRoom = errors.New("only staff members can create rooms")
	ErrCreateRoom          = errors.New(
		"something went wrong while trying to create a new room. Please, try again. If the error persists, contact the development team",
	)
)
