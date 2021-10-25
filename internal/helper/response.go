package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

/*
func JSONOKResponse(w http.ResponseWriter) {
	var ok struct {
		Result string `json:"result"`
	}

	ok.Result = "OK"

	JSONResponse(w, ok)
}




func NoContentResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func CreatedResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		goerrors.Log().WithError(err).Error("cannot write response")
	}
}
*/

func JSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("cannot write response")
	}
	log.Printf("status %d: %s", http.StatusOK, http.StatusText(http.StatusOK))

}

func ErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	if httpErr, ok := err.(*HTTPError); ok {
		log.Printf("status %d: %s", httpErr.Status, http.StatusText(httpErr.Status))
		w.WriteHeader(httpErr.Status)
		_ = json.NewEncoder(w).Encode(Error{
			Error: httpErr.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	log.Printf("status %d: %s", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	_ = json.NewEncoder(w).Encode(Error{
		Error: defaultInternalServerErrorMessage,
	})
}
