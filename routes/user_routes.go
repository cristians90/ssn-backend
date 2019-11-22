package routes

import (
	"github.com/go-chi/chi"
	"net/http"
	"ssn-backend/handlers"
)

func GetUserRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/{userID}", handlers.GetUserHandler)
	r.Put("/", handlers.PutUserHandler)

	return r
}
