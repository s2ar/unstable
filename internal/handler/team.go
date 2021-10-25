package handler

import (
	"crypto/rand"
	"math/big"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/s2ar/unstable/internal/application"
	"github.com/s2ar/unstable/internal/helper"
)

func GenRouting(r *mux.Router) {
	r.Path("/team/top").Methods(http.MethodGet).HandlerFunc(handleTeamGetInfo)
}

func handleTeamGetInfo(w http.ResponseWriter, r *http.Request) {

	status := imitationOfError(w, r)
	if status {
		return
	}

	app, err := application.GetAppFromRequest(r)
	if err != nil {
		helper.ErrorResponse(w, r, err)
		return
	}

	serviceOpendota := app.ServiceOpendota()

	response, err := serviceOpendota.GetTopTeam()
	if err != nil {
		helper.ErrorResponse(w, r, err)
		return
	}

	helper.JSONResponse(w, response)
}

func imitationOfError(w http.ResponseWriter, r *http.Request) bool {
	randNum, err := rand.Int(rand.Reader, big.NewInt(10))
	if err != nil {
		helper.ErrorResponse(w, r, err)
		return true
	}

	if randNum.Int64() < 1 {
		time.Sleep(3 * time.Second)
		helper.ErrorResponse(w, r, helper.NewHTTPError(nil, http.StatusRequestTimeout, "request timeout"))
		return true
	} else if randNum.Int64() < 3 {
		helper.ErrorResponse(w, r, nil)
		return true
	}
	return false
}
