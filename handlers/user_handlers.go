package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"ssnbackend/repository/user"
	"ssnbackend/utils/handler"
	"strconv"

	"github.com/go-chi/chi"
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

func PostAvatarHandler(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, h, err := r.FormFile("imageFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		handler.WriteEmptyResponse(w, http.StatusBadRequest)
		return
	}

	defer file.Close()
	fmt.Printf("File Size: %+v\n", h.Size)
	fmt.Printf("MIME Header: %+v\n", h.Header)

	contentType := h.Header.Get("Content-Type")

	fmt.Printf("Content Type: %+v\n", contentType)

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		handler.WriteEmptyResponse(w, http.StatusInternalServerError)
		return
	}

	userContext := handler.GetUserFromContext(r)

	err = user.SetAvatar(fileBytes, contentType, userContext.IDUser)

	if err == nil {
		handler.WriteEmptyResponse(w, http.StatusOK)
	} else {
		handler.WriteEmptyResponse(w, http.StatusInternalServerError)
	}
}

func GetAvatarHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	id, err := strconv.ParseUint(userID, 10, 64)

	if err != nil {
		handler.WriteJSONErrorResponse(w, "Invalid URL parameter", http.StatusBadRequest)
		return
	}

	foundAvatar, err := user.GetAvatar(id)

	if err != nil {
		handler.WriteJSONErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", foundAvatar.BinaryContentType)
	w.Header().Set("Content-Length", strconv.Itoa(len(foundAvatar.BinaryImage)))
	w.WriteHeader(http.StatusOK)
	w.Write(foundAvatar.BinaryImage)
}

func PutUserHandler(w http.ResponseWriter, r *http.Request) {

	//serContext := handler.GetUserFromContext(r)

	w.Write([]byte(fmt.Sprintf("PUT USER HANDLER!")))
}
