package roomdb

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/gclenz/tinybookingapi/internal/app/domain/room"
)

type RoomRepository struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) room.Repository {
	return &RoomRepository{
		db: db,
	}
}

// Create implements room.Repository.
func (rr *RoomRepository) Create(room *room.Room, ctx context.Context) error {
	_, err := rr.db.ExecContext(
		ctx,
		InsertRoomQuery,
		&room.ID,
		&room.Name,
		&room.SingleBedCount,
		&room.DoubleBedCount,
		&room.GuestsLimit,
		&room.ArePetsAllowed,
		&room.CreatedAt,
		&room.UpdatedAt,
	)

	if err != nil {
		slog.Error("RoomRepository(Create) error:", err)
		return err
	}

	return nil
}

// FindByID implements room.Repository.
func (rr *RoomRepository) FindByID(roomID string, ctx context.Context) (*room.Room, error) {
	row := rr.db.QueryRowContext(
		ctx,
		SelectRoomByIDQuery,
		roomID,
	)
	err := row.Err()
	if err != nil {
		slog.Error("RoomRepository(FindByID) error:", err)
		return nil, err
	}

	var room *room.Room

	err = row.Scan(room)
	if err != nil {
		slog.Error("RoomRepository(FindByID) error:", err)
		return nil, err
	}

	return room, nil
}
