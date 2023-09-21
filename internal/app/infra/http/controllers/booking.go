package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gclenz/tinybookingapi/internal/app/domain/booking"
	"github.com/gclenz/tinybookingapi/internal/app/infra/http/middlewares"
)

type CreateBookingRequest struct {
	RoomID  string    `json:"roomId"`
	StartOn time.Time `json:"startOn"`
	EndOn   time.Time `json:"endOn"`
}

type BookingController struct {
	createBooking *booking.CreateBooking
}

func (bc *BookingController) CreateBooking(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var bookingData CreateBookingRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&bookingData)
	if err != nil {
		slog.Error("BookingController(Create) error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx := r.Context()
	userID, ok := ctx.Value(middlewares.ContextUserID).(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, err = bc.createBooking.Execute(userID, bookingData.RoomID, bookingData.StartOn, bookingData.EndOn, ctx)
	if err != nil {
		slog.Info(err.Error())
		switch {
		case errors.Is(booking.ErrOverlappingBooking, err),
			errors.Is(booking.ErrCreateBooking, err),
			errors.Is(booking.ErrEndDateBeforeStartAtDate, err),
			errors.Is(booking.ErrBookingInThePast, err):
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(fmt.Sprintf(`{"message": "%s."}`, err)))
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Something went wrong. The booking could not be created."}`))
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func NewCreateBookingController(createBooking *booking.CreateBooking) *BookingController {
	return &BookingController{
		createBooking: createBooking,
	}
}
