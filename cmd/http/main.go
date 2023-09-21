package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gclenz/tinybookingapi/internal/app/infra/database"
	"github.com/gclenz/tinybookingapi/internal/app/infra/http/middlewares"
	"github.com/gclenz/tinybookingapi/internal/app/infra/http/routers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("error loading environment variables:", err)
		os.Exit(1)
	}
	db := database.GetDatabaseConnection()
	userRouter := routers.NewUserRouter(db)
	roomRouter := routers.NewRoomRouter(db)
	bookingRouter := routers.NewBookingRouter(db)

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(middleware.Logger)
	r.Use(middlewares.Authentication)
	r.Mount("/users", userRouter)
	r.Mount("/rooms", roomRouter)
	r.Mount("/bookings", bookingRouter)
	fmt.Printf("Server startAted at: %v", time.Now().UTC())
	http.ListenAndServe(":8080", r)
}
