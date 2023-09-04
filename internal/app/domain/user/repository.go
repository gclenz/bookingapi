package user

type Repository interface {
	Create(*User) error
	FindByID(userID string) (*User, error)
}
