package routers

import (
	"database/sql"
	"os"

	"github.com/gclenz/tinybookingapi/internal/app/domain/user"
	"github.com/gclenz/tinybookingapi/internal/app/infra/database/userdb"
	"github.com/gclenz/tinybookingapi/internal/app/infra/email"
	"github.com/gclenz/tinybookingapi/internal/app/infra/http/controllers"
	"github.com/gclenz/tinybookingapi/internal/app/infra/http/utils"
	"github.com/go-chi/chi/v5"
)

func NewUserRouter(db *sql.DB) chi.Router {
	mux := chi.NewMux()

	ur := userdb.NewUserRepository(db)

	jwtHandler := utils.JWTHandler{
		Secret: os.Getenv("JWT_SECRET"),
	}
	emailService := email.NewEmailService(
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"),
		os.Getenv("SMTP_HOST"),
	)
	createUser := user.NewCreate(ur)
	authenticate := user.NewAuthenticate(ur, &jwtHandler)
	requestAuthenticationCode := user.NewRequestAuthenticationCode(emailService, ur)
	userController := controllers.NewUserController(createUser, requestAuthenticationCode, authenticate, &jwtHandler)

	router := mux.Group(func(r chi.Router) {
		r.Post("/", userController.CreateUser)
		r.Post("/authenticate", userController.AuthenticateUser)
		r.Post("/code", userController.RequestAuthenticationCode)
	})
	return router
}
