package handler

import (
	"context"
	"net/http"
	"ssnbackend/utils/appcontext"
)

const userContextKey = "userContext"

func SetUserInContext(r *http.Request, userContext appcontext.UserContext) context.Context {
	ctx := context.WithValue(r.Context(), userContextKey, userContext)
	return ctx
}

func GetUserFromContext(r *http.Request) appcontext.UserContext {
	value, _ := r.Context().Value(userContextKey).(appcontext.UserContext)
	return value
}
