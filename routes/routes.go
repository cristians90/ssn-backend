package routes

import (
	"github.com/go-chi/chi"
	"net/http"
	"ssn-backend/middlewares"
)

func GetRoutes() http.Handler {
	r := chi.NewRouter()

	r.Route("/api/auth", func(r chi.Router) {
		r.Mount("/", GetAuthRoutes())
	})

	r.Route("/api", func(r chi.Router) {
		r.Use(middlewares.TokenValidationMiddleware)
		r.Mount("/user", GetUserRoutes())
		r.Mount("/post", GetPostRoutes())
	})

	return r
}
