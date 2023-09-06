package routers

import (
	"database/sql"

	"github.com/gclenz/tinybookingapi/internal/app/domain/user"
	"github.com/gclenz/tinybookingapi/internal/app/infra/database/userdb"
	"github.com/gclenz/tinybookingapi/internal/app/infra/http/controllers"
	"github.com/go-chi/chi/v5"
)

func NewUserRouter(db *sql.DB) chi.Router {
	mux := chi.NewMux()

	ur := userdb.NewUserRepository(db)

	createUser := user.NewCreateUser(ur)
	createUC := controllers.NewCreateUserController(createUser)

	router := mux.Group(func(r chi.Router) {
		r.Post("/", createUC.CreateUser)
	})
	return router
}
