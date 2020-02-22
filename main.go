package main

import (
	"log"
	"net/http"

	"github.com/rusq/osenv"
)

const (
	root              = "/"
	apiRoot           = "/api/v1"
	apiRubbishRecycle = apiRoot + "/rr/"
	apiRRExt          = apiRoot + "/rrext/"
)

var port = osenv.String("PORT", "8080")

func main() {
	http.HandleFunc(root, rootHandler)
	http.HandleFunc(apiRubbishRecycle, rrHandler)
	http.HandleFunc(apiRRExt, rrExtHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
