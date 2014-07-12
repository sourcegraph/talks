package client

import "net/http"

// START OMIT

func New(c *http.Client) *Client {
	return &Client{Repositories: &repositoriesClient{c}}
}

type Client struct{ Repositories RepositoriesService }

// END OMIT
