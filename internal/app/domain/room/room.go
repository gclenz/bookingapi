package room

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	SingleBedCount int       `json:"singleBedCount"`
	DoubleBedCount int       `json:"doubleBedCount"`
	GuestsLimit    int       `json:"guestsLimit"`
	PetFriendly    bool      `json:"petFriendly"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func NewRoom(
	name string,
	singleBedCount int,
	doubleBedCount int,
	guestsLimit int,
	petFriendly bool,
) *Room {
	now := time.Now()
	return &Room{
		ID:             uuid.NewString(),
		Name:           name,
		SingleBedCount: singleBedCount,
		DoubleBedCount: doubleBedCount,
		GuestsLimit:    guestsLimit,
		PetFriendly:    petFriendly,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}

func BuildRoom(
	id string,
	name string,
	singleBedCount int,
	doubleBedCount int,
	guestsLimit int,
	petFriendly bool,
	createdAt time.Time,
	updatedAt time.Time,
) *Room {
	return &Room{
		ID:             id,
		Name:           name,
		SingleBedCount: singleBedCount,
		DoubleBedCount: doubleBedCount,
		GuestsLimit:    guestsLimit,
		PetFriendly:    petFriendly,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}
}
