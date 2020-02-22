package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rusq/aklrubbish/aklapi"
)

const dateFmt = "2006-01-02"

type rrResponse struct {
	Rubbish string `json:"rubbish,omitempty"`
	Recycle string `json:"recycle,omitempty"`
	Address string `json:"address,omitempty"`
	Error   string `json:"error,omitempty"`
}

func respond(w http.ResponseWriter, data interface{}, code int) {
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(b)
}

func rubbish(r *http.Request) (*aklapi.CollectionDayDetailResult, error) {
	addr := r.FormValue("addr")
	if addr == "" {
		return nil, errors.New(http.StatusText(http.StatusBadRequest))
	}
	return aklapi.CollectionDayDetail(addr)
}

func rrHandler(w http.ResponseWriter, r *http.Request) {
	res, err := rubbish(r)
	if err != nil {
		respond(w, &rrResponse{Error: err.Error()}, http.StatusBadRequest)
		return
	}
	resp := rrResponse{
		Recycle: res.NextRecycle().Format(dateFmt),
		Rubbish: res.NextRubbish().Format(dateFmt),
		Address: res.Address.Address,
	}
	respond(w, resp, http.StatusOK)
}

func rrExtHandler(w http.ResponseWriter, r *http.Request) {
	res, err := rubbish(r)
	if err != nil {
		respond(w, rrResponse{Error: err.Error()}, http.StatusBadRequest)
		return
	}
	respond(w, res, http.StatusOK)
}
