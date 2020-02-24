package main

import (
	"html/template"
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
var tmpl = template.Must(template.New("index.html").Parse(rootHTML))

func main() {
	http.HandleFunc(root, rootHandler)
	http.HandleFunc(apiRubbishRecycle, rrHandler)
	http.HandleFunc(apiRRExt, rrExtHandler)

	log.Printf("Listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
