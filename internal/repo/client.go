package repo

import (
	"github.com/adhityapp/go-starterkit/bootstrap"
)

type RepoClient struct {
	DB bootstrap.Container
}

func Repo(DB bootstrap.Container) RepoClient {
	return RepoClient{
		DB: DB,
	}
}
