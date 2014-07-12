package datastore

import "sourcegraph.com/sourcegraph/srcgraph/client"

type DBHandle struct{}

// dummy
func (_ DBHandle) Query(v interface{}, sql string, args ...interface{}) error {}

// START OMIT

func New() *DataStore {
	return &DataStore{&reposStore{}}
}

type DataStore struct {
	Repositories client.RepositoriesService // reuse interface
}

type reposStore struct{ dbh DBHandle }

func (s *reposStore) Get(repo string, opt *client.RepoGetOptions) (*client.Repo, client.Response, error) {
	var repo *client.Repo // reuse Repo type
	if err := s.dbh.Query(&repo, "SELECT * FROM repo WHERE uri=$1;", repo); err != nil {
		return nil, nil, err
	}
	if opt.Stat { // handle params
		repo.Stats = s.getStats(repo)
	}
	return repo, nil, nil
}

// END OMIT

// dummy
func (s *reposStore) getStats(repo string) interface{} { return nil }
