package service

import (
	"github.com/adhityapp/go-starterkit/internal/repo"
)

type ServiceClient struct {
	repo repo.RepoClient
}

func Service(repo repo.RepoClient) ServiceClient {
	return ServiceClient{
		repo: repo,
	}
}
