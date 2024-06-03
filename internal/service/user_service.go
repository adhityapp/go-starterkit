package service

import (
	"context"

	"github.com/adhityapp/go-starterkit/internal/repo"
)

type ServiceClient struct {
	repo repo.RepoClient
}

func UserService(repo repo.RepoClient) ServiceClient {
	return ServiceClient{
		repo: repo,
	}
}

func (us ServiceClient) GetUserService(c context.Context) ([]repo.UserModel, error) {
	return us.repo.GetUser(c)
}
