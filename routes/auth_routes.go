package routes

import (
	"github.com/go-chi/chi"
	"net/http"
	"ssnbackend/handlers"
)

func GetAuthRoutes() http.Handler {
	r := chi.NewRouter()

	r.Post("/login", handlers.PostLogInHandler)
	r.Post("/signin", handlers.PostSignInHandler)
	r.Post("/refreshtoken", handlers.PostRefreshTokenHandler)

	return r
}
