package room

import "context"

type Repository interface {
	Create(room *Room, ctx context.Context) error
	FindByID(roomID string, ctx context.Context) (*Room, error)
}
