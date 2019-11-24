package routes

import (
	"github.com/go-chi/chi"
	"net/http"
	"ssnbackend/handlers"
)

func GetUserRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/{userID}", handlers.GetUserHandler)
	r.Put("/", handlers.PutUserHandler)
	r.Post("/avatar", handlers.PostAvatarHandler)
	r.Get("/{userID}/avatar", handlers.GetAvatarHandler)

	return r
}
