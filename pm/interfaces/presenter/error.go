package presenter

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/onituka/agile-project-management/project-management/apperrors"
)

type httpError struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"error"`
}

func (e *httpError) Error() string {
	b, err := json.Marshal(&e)
	if err != nil {
		log.Println(err)
	}

	return string(b)
}

func ErrorJSON(w http.ResponseWriter, err error) {
	appErr := apperrors.AsAppError(err)

	httpErr := &httpError{
		StatusCode:   appErr.StatusCode(),
		ErrorMessage: appErr.Error(),
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(httpErr.StatusCode)
	if err = json.NewEncoder(w).Encode(httpErr); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
