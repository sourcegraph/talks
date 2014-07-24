import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/go-querystring/query"
	"github.com/gorilla/schema"
	"github.com/sourcegraph/thesrc"
	"github.com/sourcegraph/thesrc/router"
	"github.com/sqs/mux"
)

// START POSTS_INTERFACE OMIT
type PostsService interface {
	Get(id int) (*Post, error)

	List(opt *PostListOptions) ([]*Post, error)

	Submit(post *Post) (created bool, err error)
}

// END POSTS_INTERFACE OMIT

// START APICLIENT OMIT
type postsClient struct{ baseURL string }

func (c *postsClient) Get(id int) (*Post, error) {
	url, _ := s.client.url(router.Post, map[string]string{"ID": strconv.Itoa(id)}, nil)

	request, _ := s.client.NewRequest("GET", url.String(), nil)

	response, _ := s.client.Do(request)

	var post *Post
	json.NewDecoder(response.Body).Decode(&post)

	return post, nil
}

// END APICLIENT OMIT

// START DATASTORE OMIT

type postsStore struct{ db *sql.DB }

func (s *postsStore) Get(id int) (*Post, error) {
	var post *Post
	s.db.Select(&post, "SELECT * FROM post WHERE id=$1", id)
	return post, nil
}

// END DATASTORE OMIT

// START APP_HTTP_HANDLER OMIT
// Frontend server HTTP handler
func servePost(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.Atoi(mux.Vars(r)["ID"])

	// Push API-related logic from frontend HTTP handlers into API Client // HL
	post, _ := apiclient.Posts.Get(id) // HL

	return renderTemplate(w, r, "posts/show.html", http.StatusOK, struct {
		Post *thesrc.Post
	}{
		Post: post,
	})
}

// END APP_HTTP_HANDLER OMIT

// START API_HTTP_HANDLER OMIT
// API server HTTP handler
func servePost(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.Atoi(mux.Vars(r)["ID"])

	// Push Datastore-related logic from API HTTP handlers into Datastore // HL
	post, _ := store.Posts.Get(id) // HL

	return writeJSON(w, post)
}

// END API_HTTP_HANDLER OMIT

// START ROUTES OMIT
// package router

const (
	Post = "post"
	// ...other routes omitted
)

func API() *mux.Router {
	m := mux.NewRouter()
	m.Path("/post").Methods("GET").Name(Post)
	// ...other routes omitted
	return m
}

// END ROUTES OMIT

// START API_ROUTER OMIT

// package api

func Handler() *mux.Router {
	apiRouter := router.App()
	apiRouter.Get(router.Post).Handler(handler(servePost))
}

// END API_ROUTER OMIT

// START API_CLIENT_ROUTE_GEN OMIT
type postsClient struct{ baseURL string }

func (c *postsClient) Get(id int) (*Post, error) {

	// Generate URL from router, no fragile hardcoded strings // HL
	url, _ := s.client.url(router.Post, map[string]string{"ID": strconv.Itoa(id)}, nil) // HL

	request, _ := s.client.NewRequest("GET", url.String(), nil)

	response, _ := s.client.Do(request)

	var post *Post
	json.NewDecoder(response.Body).Decode(&post)

	return post, nil
}

// END API_CLIENT_ROUTE_GEN OMIT

// START PARAM_STRUCT OMIT
type PostListOptions struct {
	PerPage  int
	Page     int
	CodeOnly bool
}

// END PARAM_STRUCT OMIT
var d = schema.NewDecoder()

// START PARAM_STRUCT_DECODE OMIT
// URL-decode with github.com/gorilla/schema

func servePosts(w http.ResponseWriter, r *http.Request) error {
	var opt thesrc.PostListOptions
	err := schemaDecoder.Decode(&opt, r.URL.Query()) // HL
	// ...
}

// END PARAM_STRUCT_DECODE OMIT

// START PARAM_STRUCT_ENCODE OMIT

// URL-encode with github.com/google/go-querystring/query

func (c *postsClient) List(opt *PostListOptions) ([]*Post, error) {
	var url string
	// ...
	qs, err := query.Values(opt)
	resp, err := http.Get(c.baseURL + url.String() + "?" + qs.Encode()) // HL
	// ...
}

// END PARAM_STRUCT_ENCODE OMIT

// START POSTS_MOCK OMIT

type MockPostsService struct { // Implements PostsService interface
	Get_    func(id int) (*Post, error)
	List_   func(opt *PostListOptions) ([]*Post, error)
	Submit_ func(post *Post) (bool, error)
}

func (s *MockPostsService) Get(id int) (*Post, error) {
	if s.Get_ == nil {
		return nil, nil
	}
	return s.Get_(id)
}

// ...

// END POSTS_MOCK OMIT
