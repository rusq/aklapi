package main

import (
	"log"
	"net/http"
)

const (
	apiRoot           = "/api/v1"
	apiRubbishRecycle = apiRoot + "/rr/"
	apiRRExt          = apiRoot + "/rrext/"
)

func main() {
	http.HandleFunc(apiRubbishRecycle, rrHandler)
	http.HandleFunc(apiRRExt, rrExtHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
