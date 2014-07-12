package frontend

import (
	"net/http"

	"github.com/sourcegraph/talks/google-io-2014/part1/client"
	"github.com/sqs/mux"
	"github.com/sqs/schema"
)

// START OMIT
var apiclient = client.New(http.DefaultClient)

// use github.com/gorilla/schema to decode querystrings to param structs
var schemaDecoder = schema.NewDecoder()

func serveRepo(w http.ResponseWriter, r *http.Request) error {
	routeVars := mux.Vars(r)
	var opt client.RepoGetOptions // reuse parameter struct // HL
	if err := schemaDecoder.Decode(&opt, r.URL.Query()); err != nil {
		return err
	}

	repo, resp, err := apiclient.Repositories.Get(routeVars["Repo"], &opt) // reuse API client // HL
	if err != nil {
		return err
	}

	return executeTemplate("repo.html", repo, resp) // reuse pagination & cache headers from API // HL
}

// END OMIT

// dummy
func executeTemplate(name string, data interface{}, resp client.Response) error { return nil }
