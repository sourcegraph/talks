package client

import "net/http"

// START IFACE OMIT
type RepositoriesService interface {
	Get(repo string, opt *RepoGetOptions) (*Repo, Response, error) // HL
	List(opt *ListOptions) ([]*Repo, Response, error)
	ListDependencies(repo string, opt *ListDependenciesOptions) ([]*Dep, Response, error) // OMIT
	// ...
}

// START OPT OMIT
type RepoGetOptions struct {
	CommitID    string `url:",omitempty"`
	Stats       bool   `url:",omitempty"`
	LastCommit  bool   `url:",omitempty"`
	ForceUpdate bool   `url:",omitempty"`
}

// END OPT OMIT
// END IFACE OMIT

// ListAuthors(repo string, opt *ListAuthorsOptions) ([]*Author, Response, error)

type Repo struct{ Stats interface{} }
type ListOptions struct{}
type ListDependenciesOptions struct{}
type ListAuthorsOptions struct{}
type Dep struct{}
type Author struct{}

type Response interface{}

// START IMPL OMIT

type repositoriesClient struct{ *http.Client }

func (c *repositoriesClient) Get(repo string, opt *RepoGetOptions) (*Repo, Response, error) {
	resp, err := c.Client.Get(makeURL(repo, opt))
	if err != nil {
		return nil, resp, err
	}
	var repo_ Repo
	err = unmarshalResponse(resp.Body, &repo_)
	return &repo_, resp, err
}

// END IMPL OMIT

func (c *repositoriesClient) List(opt *ListOptions) ([]*Repo, Response, error) {
	return nil, nil, nil
}
func (c *repositoriesClient) ListDependencies(repo string, opt *ListDependenciesOptions) ([]*Dep, Response, error) {
	return nil, nil, nil
}

func makeURL(v ...interface{}) string          { return "" }
func unmarshalResponse(v ...interface{}) error { return nil }
