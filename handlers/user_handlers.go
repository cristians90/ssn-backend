package handlers

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"ssn-backend/repository/user"
	"ssn-backend/utils/handler"
	"strconv"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	id, err := strconv.ParseUint(userID, 10, 64)

	if err != nil {
		handler.WriteJSONErrorResponse(w, "Invalid URL parameter", http.StatusBadRequest)
		return
	}

	foundUser, err := user.GetUserById(id)

	if err != nil {
		handler.WriteJSONErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	handler.WriteJSONResponse(w, foundUser.GetModelWithOutPassword(), http.StatusOK)
}

func PutUserHandler(w http.ResponseWriter, r *http.Request) {

	//serContext := handler.GetUserFromContext(r)

	w.Write([]byte(fmt.Sprintf("PUT USER HANDLER!")))
}
