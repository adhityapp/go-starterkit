package repo

import (
	"github.com/jmoiron/sqlx"
)

type RepoClient struct {
	DB sqlx.DB
}

func Repo(DB sqlx.DB) RepoClient {
	return RepoClient{
		DB: DB,
	}
}
