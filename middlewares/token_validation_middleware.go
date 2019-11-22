package middlewares

import (
	"net/http"
	"ssnbackend/utils/handler"
	"ssnbackend/utils/token"
	"strings"
)

func TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authString := r.Header.Get("Authorization")
		authString = strings.Replace(authString, "Bearer ", "", -1)

		userContext, err := token.ValidateToken(authString)

		if err != nil {
			handler.WriteJSONErrorResponse(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		ctx := handler.SetUserInContext(r, userContext)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
