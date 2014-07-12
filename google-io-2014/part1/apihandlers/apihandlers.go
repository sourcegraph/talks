package apihandlers

import (
	"net/http"

	"github.com/sourcegraph/talks/google-io-2014/part1/client"
	"github.com/sourcegraph/talks/google-io-2014/part1/datastore"
	"github.com/sqs/mux"
	"github.com/sqs/schema"
)

// use github.com/gorilla/schema to decode querystrings to param structs
var schemaDecoder = schema.NewDecoder()

// START ROUTER OMIT
func init() {
	r := client.NewAPIRouter()
	// Get existing named route and mount a handler on it
	r.Get(client.RepoRoute).Handler(handleErr(serveRepo))
	http.Handle("/api", r)
}

// END ROUTER OMIT

// START OMIT

var store = datastore.New()

func serveRepo(w http.ResponseWriter, r *http.Request) error {
	routeVars := mux.Vars(r)
	var opt client.RepoGetOptions // reuse parameter struct // HL
	if err := schemaDecoder.Decode(&opt, r.URL.Query()); err != nil {
		return err
	}

	repo, resp, err := store.Repositories.Get(routeVars["Repo"], &opt) // reuse data store // HL
	if err != nil {
		return err
	}

	// check authorization, rate limits, etc., here.

	return writeJSON(repo, resp) // reuse pagination & cache info from data store // HL
}

// END OMIT

type handleErr func(http.ResponseWriter, *http.Request) error

// dummy
func (_ handleErr) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
