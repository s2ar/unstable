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

const maxRandNum = 10

func GenRouting(r *mux.Router) {
	r.Path("/team/top").Methods(http.MethodGet).HandlerFunc(TeamGetInfo)
}

func TeamGetInfo(w http.ResponseWriter, r *http.Request) {
	randNum, err := rand.Int(rand.Reader, big.NewInt(maxRandNum))
	if err != nil {
		helper.ErrorResponse(w, r, err)
		return
	}

	status := imitationOfError(w, r, randNum.Int64())
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

func imitationOfError(w http.ResponseWriter, r *http.Request, rand int64) bool {

	if rand < 1 {
		time.Sleep(3 * time.Second)
		helper.ErrorResponse(w, r, helper.NewHTTPError(nil, http.StatusRequestTimeout, "request timeout"))
		return true
	} else if rand < 3 {
		helper.ErrorResponse(w, r, nil)
		return true
	}
	return false
}
