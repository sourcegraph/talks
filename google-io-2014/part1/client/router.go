package client

import "github.com/sqs/mux"

// START ROUTER OMIT

const RepoRoute = "repo"

func NewAPIRouter() *mux.Router {
	m := mux.NewRouter()
	// Define a named route but don't mount a handler (yet)
	m.Path("/repos/{Repo:.*}").Methods("GET").Name(RepoRoute)
	return m
}

// END ROUTER OMIT
