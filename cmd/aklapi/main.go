package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/rusq/aklapi"
	"github.com/rusq/osenv/v2"
)

const (
	banner = "Auckland Council API Server (unofficial)"

	defaultAddr = ""
	defaultPort = "8080"
)

const (
	root              = "/"
	apiHealthcheck    = "/healthcheck"
	apiRoot           = "/api/v1"
	apiAddr           = apiRoot + "/addr/"
	apiRubbishRecycle = apiRoot + "/rr/"
	apiRRExt          = apiRoot + "/rrext/"
)

func init() {
	flag.Usage = usage
}

var (
	host    = flag.String("host", osenv.Value("HOST", defaultAddr), "host to listen on")
	port    = flag.String("port", osenv.Value("PORT", defaultPort), "port to listen on")
	nocache = flag.Bool("no-cache", osenv.Value("NO_CACHE", false), "disable caching")
)

var tmpl = template.Must(template.New("index.html").Parse(rootHTML))

func main() {
	flag.Parse()
	if *port == "" {
		log.Printf("no port specified, defaulting to %s", defaultPort)
	}

	// Set the global caching flag.
	aklapi.NoCache = *nocache

	http.HandleFunc(root, rootHandler)
	http.HandleFunc(apiHealthcheck, healthcheckHandler)
	http.HandleFunc(apiAddr, addrHandler)
	http.HandleFunc(apiRubbishRecycle, rrHandler)
	http.HandleFunc(apiRRExt, rrExtHandler)

	hostport := fmt.Sprintf("%s:%s", *host, *port)
	log.Println(banner)
	log.Println("Listening on: ", hostport)
	log.Fatal(http.ListenAndServe(hostport, nil))
}

func usage() {
	fmt.Fprintln(flag.CommandLine.Output(), banner)
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options]\n", os.Args[0])
	flag.CommandLine.PrintDefaults()
}
