package handler

import (
	"context"
	"net/http"
	"ssn-backend/utils/app_context"
)

const userContextKey = "userContext"

func SetUserInContext(r *http.Request, userContext app_context.UserContext) context.Context {
	ctx := context.WithValue(r.Context(), userContextKey, userContext)
	return ctx
}

func GetUserFromContext(r *http.Request) app_context.UserContext {
	value, _ := r.Context().Value(userContextKey).(app_context.UserContext)
	return value
}
