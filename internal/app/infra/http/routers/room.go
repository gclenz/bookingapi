package routers

import (
	"database/sql"

	"github.com/gclenz/tinybookingapi/internal/app/domain/room"
	"github.com/gclenz/tinybookingapi/internal/app/infra/database/roomdb"
	"github.com/gclenz/tinybookingapi/internal/app/infra/database/userdb"
	"github.com/gclenz/tinybookingapi/internal/app/infra/http/controllers"
	"github.com/go-chi/chi/v5"
)

func NewRoomRouter(db *sql.DB) chi.Router {
	mux := chi.NewMux()

	rr := roomdb.NewRoomRepository(db)
	ur := userdb.NewUserRepository(db)

	createRoom := room.NewCreateRoom(rr, ur)
	createRC := controllers.NewCreateRoomController(createRoom)

	router := mux.Group(func(r chi.Router) {
		r.Post("/", createRC.CreateRoom)
	})
	return router
}
