package controllers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/gclenz/tinybookingapi/internal/app/domain/room"
	"github.com/gclenz/tinybookingapi/internal/app/infra/http/middlewares"
)

type CreateRoomRequest struct {
	Name           string `json:"name"`
	SingleBedCount int    `json:"singleBedCount"`
	DoubleBedCount int    `json:"doubleBedCount"`
	GuestsLimit    int    `json:"guestsLimit"`
	ArePetsAllowed bool   `json:"arePetsAllowed"`
}

type RoomController struct {
	createRoom *room.CreateRoom
}

func (rc *RoomController) CreateRoom(w http.ResponseWriter, r *http.Request) {
	defer func() {
		r.Body.Close()
	}()

	var roomReq CreateRoomRequest

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&roomReq)
	if err != nil {
		slog.Error("CreateRoom error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx := r.Context()
	userID, ok := ctx.Value(middlewares.ContextUserID).(string)
	slog.Info("userID", userID)
	if !ok {
		slog.Error("userID not found on context", userID)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = rc.createRoom.Execute(
		userID,
		roomReq.Name,
		roomReq.SingleBedCount,
		roomReq.DoubleBedCount,
		roomReq.GuestsLimit,
		roomReq.ArePetsAllowed,
		ctx,
	)
	if err != nil {
		slog.Error("CreateRoom error:", err)
		switch {
		case errors.Is(room.ErrStaffOnlyCreateRoom, err):
			w.WriteHeader(http.StatusForbidden)
			return
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Something went wrong. The room could not be created."}`))
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func NewCreateRoomController(createRoom *room.CreateRoom) *RoomController {
	return &RoomController{
		createRoom: createRoom,
	}
}
