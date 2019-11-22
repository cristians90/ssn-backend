package routes

import (
	"github.com/go-chi/chi"
	"net/http"
	"ssn-backend/handlers"
)

func GetPostRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/{postID}", handlers.GetPostWithIdHandler)
	r.Get("/", handlers.GetPostHandler)
	r.Post("/", handlers.PostPostHandler)

	return r
}
