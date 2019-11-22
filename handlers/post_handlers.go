package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"ssn-backend/repository/models"
	"ssn-backend/repository/post"
	"ssn-backend/utils/handler"
	"strconv"
	"time"
)

func GetPostHandler(w http.ResponseWriter, r *http.Request) {

	sinceDate, ok := r.URL.Query()["sinceDate"]

	if !ok || len(sinceDate[0]) < 1 {
		handler.WriteJSONErrorResponse(w, "Url Param 'sinceDate' is missing", http.StatusBadRequest)
		return
	}

	date, err := parseStringToTime(sinceDate[0])

	if err != nil {
		handler.WriteJSONErrorResponse(w, "Url Param 'sinceDate' is invalid", http.StatusBadRequest)
		return
	}

	limitResults, ok := r.URL.Query()["limitResults"]

	if !ok || len(limitResults[0]) < 1 {
		handler.WriteJSONErrorResponse(w, "Url Param 'limitResults' is missing", http.StatusBadRequest)
		return
	}

	lr, err := strconv.Atoi(limitResults[0])

	if err != nil {
		handler.WriteJSONErrorResponse(w, "Url Param 'limitResults' is invalid", http.StatusBadRequest)
		return
	}

	offsetResults, ok := r.URL.Query()["offsetResults"]

	if !ok || len(offsetResults[0]) < 1 {
		handler.WriteJSONErrorResponse(w, "Url Param 'offsetResults' is missing", http.StatusBadRequest)
		return
	}

	or, err := strconv.Atoi(offsetResults[0])

	if err != nil {
		handler.WriteJSONErrorResponse(w, "Url Param 'offsetResults' is invalid", http.StatusBadRequest)
		return
	}

	posts, err := post.GetPosts(date, lr, or)

	if err != nil {
		handler.WriteJSONErrorResponse(w, "Fatal error", http.StatusInternalServerError)
		return
	}

	handler.WriteJSONResponse(w, posts, http.StatusOK)
}

func GetPostWithIdHandler(w http.ResponseWriter, r *http.Request) {

}

func PostPostHandler(w http.ResponseWriter, r *http.Request) {
	type jsonModel struct {
		Content string `json:"content"`
	}

	var json jsonModel

	err := handler.DecodeJSONBody(*r, &json)
	if err != nil {
		handler.WriteJSONErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	userContext := handler.GetUserFromContext(r)

	postObj := models.PostModel{
		Content:   json.Content,
		CreatedBy: userContext.IDUser,
	}

	err = post.InsertPost(postObj)

	if err != nil {
		handler.WriteJSONErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	handler.WriteEmptyResponse(w, http.StatusCreated)
}

func parseStringToTime(date string) (time.Time, error) {
	var result time.Time

	re := regexp.MustCompile(`\d{14}$`)

	if len(re.FindStringIndex(date)) == 0 {
		return result, errors.New("invalid string")
	}

	year := date[0:4]
	month := date[4:6]
	day := date[6:8]
	hour := date[8:10]
	minute := date[10:12]
	second := date[12:14]

	datetime := fmt.Sprintf("%s-%s-%s %s:%s:%s", year, month, day, hour, minute, second)

	result, err := time.Parse("2006-01-02 15:04:05", datetime)

	if err != nil {
		return result, errors.New("invalid date")
	}

	return result, nil
}
