package room

type Repository interface {
	Create(room *Room) error
	FindByID(roomID string) (*Room, error)
}
