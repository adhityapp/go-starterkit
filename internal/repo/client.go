package repo

import (
	"github.com/adhityapp/go-starterkit/bootstrap"
)

type RepoClient struct {
	Container *bootstrap.Container
}

func Repo(Container *bootstrap.Container) RepoClient {
	return RepoClient{
		Container: Container,
	}
}
