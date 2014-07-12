package external_use

import (
	"net/http"

	"github.com/sourcegraph/talks/google-io-2014/part1/client"
)

func getRepo() {
	// START OMIT
	c := client.New(http.DefaultClient)
	opt := &client.RepoGetOptions{CommitID: "2f3cf5"}
	repo, _, err := c.Repositories.Get("github.com/gorilla/mux", opt)
	// END OMIT
	_ = repo
	_ = err
}
