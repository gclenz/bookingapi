package controllers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gclenz/tinybookingapi/internal/app/domain/room"
	"github.com/gclenz/tinybookingapi/internal/app/infra/http/middlewares"
	"github.com/gclenz/tinybookingapi/internal/app/infra/http/utils"
)

type CreateRoomRequest struct {
	Name           string `json:"name"`
	SingleBedCount int    `json:"singleBedCount"`
	DoubleBedCount int    `json:"doubleBedCount"`
	GuestsLimit    int    `json:"guestsLimit"`
	PetFriendly    bool   `json:"petFriendly"`
}

type RoomController struct {
	createRoom *room.CreateRoom
}

func (rc *RoomController) CreateRoom(w http.ResponseWriter, r *http.Request) {
	defer func() {
		r.Body.Close()
	}()

	var roomReq CreateRoomRequest

	err := utils.ParseJSON(&roomReq, w, r)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		var mr *utils.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Message, mr.Status)
			return
		}
		slog.Error("CreateRoom error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Something went wrong.`))
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
		roomReq.PetFriendly,
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
