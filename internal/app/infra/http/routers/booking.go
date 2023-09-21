package routers

import (
	"database/sql"

	"github.com/gclenz/tinybookingapi/internal/app/domain/booking"
	"github.com/gclenz/tinybookingapi/internal/app/infra/database/bookingdb"
	"github.com/gclenz/tinybookingapi/internal/app/infra/database/roomdb"
	"github.com/gclenz/tinybookingapi/internal/app/infra/database/userdb"
	"github.com/gclenz/tinybookingapi/internal/app/infra/http/controllers"
	"github.com/go-chi/chi/v5"
)

func NewBookingRouter(db *sql.DB) chi.Router {
	mux := chi.NewMux()

	br := bookingdb.NewBookingRepository(db)
	rr := roomdb.NewRoomRepository(db)
	ur := userdb.NewUserRepository(db)

	createBooking := booking.NewCreateBooking(br, rr, ur)

	createBC := controllers.NewCreateBookingController(createBooking)

	router := mux.Group(func(r chi.Router) {
		r.Post("/", createBC.CreateBooking)
	})
	return router
}
