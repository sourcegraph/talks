package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/google/go-querystring/query"
	"github.com/sqs/mux"
	"github.com/sqs/schema"
)

// START SVC OMIT
type RepositoriesService interface {
	Get(name string) (*Repo, error)
	List() ([]*Repo, error)
	Search(opt *SearchOptions) ([]*Repo, error)
	// ...
}

// END SVC OMIT

type Repo struct {
	Name     string
	CloneURL string
}

// START SEARCH OPTIONS OMIT
// options for method: Search(opt *SearchOptions) ([]*Repo, error) // HL
type SearchOptions struct {
	Owner    string
	Language string
}

// END SEARCH OPTIONS OMIT

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// START API GET OMIT
var repoDataStore RepositoriesService = &repoStore{}

func handleRepoGet(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["Name"]
	repo, _ := repoDataStore.Get(name)
	b, _ := json.Marshal(repo)
	w.Write(b)
}

// END API GET OMIT

// START API LIST OMIT

func handleRepoList(w http.ResponseWriter, r *http.Request) {
	repos, _ := repoDataStore.List()
	b, _ := json.Marshal(repos)
	w.Write(b)
}

// END API GET OMIT

// START API SEARCH OMIT
var d = schema.NewDecoder()

func handleRepoSearch(w http.ResponseWriter, r *http.Request) {
	var opt SearchOptions
	d.Decode(&opt, r.URL.Query()) // decode querystring with github.com/gorilla/schema // HL
	// ...
	// END API SEARCH OMIT
	repos, _ := repoDataStore.Search(&opt)
	b, _ := json.Marshal(repos)
	w.Write(b)
}

// START API ROUTER OMIT
const (
	RepoGetRoute    = "repo"
	RepoListRoute   = "repo.list"
	RepoSearchRoute = "repo.search" // OMIT
)

func NewAPIRouter() *mux.Router {
	m := mux.NewRouter()
	// define the routes // HL
	m.Path("/api/repos/search").Name(RepoSearchRoute) // OMIT
	m.Path("/api/repos/{Name:.*}").Name(RepoGetRoute)
	m.Path("/api/repos").Name(RepoListRoute)
	return m
}

func init() {
	m := NewAPIRouter()
	// mount handlers // HL
	m.Get(RepoGetRoute).HandlerFunc(handleRepoGet)
	m.Get(RepoListRoute).HandlerFunc(handleRepoList)
	m.Get(RepoSearchRoute).HandlerFunc(handleRepoSearch)
	http.Handle("/api/", m)
}

// END API ROUTER OMIT

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// START FRONTEND OMIT
var repoAPIClient RepositoriesService = &repoClient{"http://localhost:7777"}

func handleRepoPage(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["Name"]
	repo, _ := repoAPIClient.Get(name) // HL
	fmt.Fprintf(w, "<h1>%s</h1><p>Clone URL: %s</p>", repo.Name, repo.CloneURL)
}

// END FRONTEND OMIT

func handleRepoSearchPage(w http.ResponseWriter, r *http.Request) {
	var opt SearchOptions
	d.Decode(&opt, r.URL.Query()) // decode querystring with github.com/gorilla/schema // HL
	repos, _ := repoAPIClient.Search(&opt)
	fmt.Fprintf(w, "<h1>Search: %+v</h1>", opt)
	for _, repo := range repos {
		fmt.Fprintf(w, `<p>%s (<a href="%s">%s</a>)</p>`, repo.Name, repo.CloneURL, repo.CloneURL)
	}
}

func init() {
	m := mux.NewRouter()
	m.Path("/repos/search").HandlerFunc(handleRepoSearchPage)
	m.Path("/repos/{Name:.*}").HandlerFunc(handleRepoPage)
	http.Handle("/", m)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// START CLIENT OMIT
type repoClient struct{ baseURL string }

func (s *repoClient) Get(name string) (*Repo, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/repos/%s", s.baseURL, name))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var repo Repo
	return &repo, json.NewDecoder(resp.Body).Decode(&repo)
}

// END CLIENT OMIT

// START CLIENT LIST OMIT

var apiRouter = NewAPIRouter()

func (s *repoClient) List() ([]*Repo, error) {
	url, _ := apiRouter.Get(RepoListRoute).URL() // HL
	resp, err := http.Get(s.baseURL + url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var repos []*Repo
	return repos, json.NewDecoder(resp.Body).Decode(&repos)
}

// END CLIENT LIST OMIT

// START CLIENT SEARCH OMIT

func (s *repoClient) Search(opt *SearchOptions) ([]*Repo, error) {
	url, _ := apiRouter.Get(RepoSearchRoute).URL()
	q, _ := query.Values(opt) // encode querystring with github.com/google/go-querystring/query // HL
	resp, err := http.Get(s.baseURL + url.String() + "?" + q.Encode())
	// ...
	if err != nil { // OMIT
		return nil, err // OMIT
	} // OMIT
	defer resp.Body.Close() // OMIT
	// OMIT
	var repos []*Repo                                       // OMIT
	return repos, json.NewDecoder(resp.Body).Decode(&repos) // OMIT
}

// END CLIENT SEARCH OMIT

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// START STORE OMIT
type repoStore struct{ db *db }

func (s *repoStore) Get(name string) (*Repo, error) {
	var repo *Repo
	return repo, s.db.Select(&repo, "SELECT * FROM repo WHERE name=$1", name)
}

// END STORE OMIT

func (s *repoStore) List() ([]*Repo, error) { return nil, nil }

func (s *repoStore) Search(opt *SearchOptions) ([]*Repo, error) {
	log.Printf("repo search options: %+v", opt)
	return []*Repo{{"myrepo", "git://github.com/foo/myrepo.git"}, {"mux", "git://github.com/gorilla/mux.git"}}, nil
}

type db struct{}

func (_ *db) Select(v interface{}, sql string, args ...interface{}) error {
	if repo, ok := v.(**Repo); ok {
		name, _ := args[0].(string)
		*repo = &Repo{filepath.Base(name), "git://" + name + ".git"}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	bind := ":7777"
	log.Printf("Listening on %s", bind)
	log.Println(http.ListenAndServe(bind, nil))
}
