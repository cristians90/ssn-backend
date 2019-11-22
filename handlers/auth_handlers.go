package handlers

import (
	"net/http"
	"ssnbackend/repository/models"
	"ssnbackend/repository/user"
	"ssnbackend/utils/appcontext"
	"ssnbackend/utils/handler"
	"ssnbackend/utils/token"
)

func PostSignInHandler(w http.ResponseWriter, r *http.Request) {
	type jsonModel struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		LastName string `json:"lastName"`
		Password string `json:"password"`
	}

	var json jsonModel

	err := handler.DecodeJSONBody(*r, &json)
	if err != nil {
		handler.WriteJSONErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	userModel := models.UserModel{Username: json.Username, Name: json.Name, LastName: json.LastName, Email: json.Email, Password: json.Password}

	err = user.InsertUser(userModel)

	if err != nil {
		handler.WriteJSONErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	handler.WriteEmptyResponse(w, http.StatusOK)
}

func PostLogInHandler(w http.ResponseWriter, r *http.Request) {
	type jsonModel struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var json jsonModel

	err := handler.DecodeJSONBody(*r, &json)
	if err != nil {
		handler.WriteJSONErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	userInDb, err := user.GetUserByUsername(json.Username)

	if err != nil {
		handler.WriteJSONErrorResponse(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if userInDb.Password != json.Password {
		handler.WriteJSONErrorResponse(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	tokenPayload := appcontext.UserContext{IDUser: userInDb.ID, UserName: userInDb.Username}

	stringToken, _ := token.GenerateToken(tokenPayload)
	stringRefreshToken, _ := token.GenerateRefreshToken(tokenPayload)

	response := tokenAndRefreshTokenResponse{Token: stringToken, RefreshToken: stringRefreshToken}

	handler.WriteJSONResponse(w, response, http.StatusOK)
}

func PostRefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	type jsonModel struct {
		RefreshToken string `json:"refreshToken"`
	}

	var json jsonModel

	err := handler.DecodeJSONBody(*r, &json)
	if err != nil {
		handler.WriteJSONErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := token.ValidateRefreshToken(json.RefreshToken)

	if err != nil {
		handler.WriteJSONErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	userInDb, err := user.GetUserById(id)

	if err != nil {
		handler.WriteJSONErrorResponse(w, "Server error", http.StatusInternalServerError)
		return
	}

	tokenPayload := appcontext.UserContext{IDUser: userInDb.ID, UserName: userInDb.Username}

	stringToken, _ := token.GenerateToken(tokenPayload)
	stringRefreshToken, _ := token.GenerateRefreshToken(tokenPayload)

	response := tokenAndRefreshTokenResponse{Token: stringToken, RefreshToken: stringRefreshToken}

	handler.WriteJSONResponse(w, response, http.StatusOK)
}

type tokenAndRefreshTokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
