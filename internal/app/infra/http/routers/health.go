package routers

import (
	"time"

	"github.com/gclenz/tinybookingapi/internal/app/infra/http/controllers"
	"github.com/go-chi/chi/v5"
)

func NewHealthRouter() chi.Router {
	mux := chi.NewMux()

	hc := controllers.NewHealthController(time.Now())

	router := mux.Group(func(r chi.Router) {
		r.Get("/", hc.Health)
	})
	return router
}
