package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/rusq/aklapi"
)

const dateFmt = "2006-01-02"

type rrResponse struct {
	Rubbish    string `json:"rubbish,omitempty"`
	Recycle    string `json:"recycle,omitempty"`
	FoodScraps string `json:"foodscraps,omitempty"`
	Address    string `json:"address,omitempty"`
	Error      string `json:"error,omitempty"`
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

func addrHandler(w http.ResponseWriter, r *http.Request) {
	addr := r.FormValue("addr")
	if addr == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	resp, err := aklapi.AddressLookup(addr)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}
	respond(w, resp, http.StatusOK)
}

func rrHandler(w http.ResponseWriter, r *http.Request) {
	res, err := rubbish(r)
	if err != nil {
		respond(w, &rrResponse{Error: err.Error()}, http.StatusBadRequest)
		return
	}
	resp := rrResponse{
		Recycle:    res.NextRecycle().Format(dateFmt),
		Rubbish:    res.NextRubbish().Format(dateFmt),
		FoodScraps: res.NextFoodScraps().Format(dateFmt),
		Address:    res.Address.Address,
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

func rootHandler(w http.ResponseWriter, r *http.Request) {
	var page = struct {
		CyberdyneLogo string
	}{
		CyberdyneLogo: cyberdynePng,
	}
	if err := tmpl.ExecuteTemplate(w, "index.html", &page); err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
