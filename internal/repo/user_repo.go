package repo

import (
	"context"
)

func (r RepoClient) GetUser(c context.Context) ([]UserModel, error) {
	var um []UserModel
	query := "select user_id, username, password, email, user_role from accounts"
	rows, err := r.Container.Dbr().QueryxContext(c, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var m UserModel
		err = rows.StructScan(&m)
		if err != nil {
			return nil, err
		}
		um = append(um, m)
	}
	return um, nil
}

func (r RepoClient) GetUserbyUsername(c context.Context, username string) (*UserModel, error) {
	var m UserModel
	query := "select user_id, username, password, email, user_role from accounts where username = ?"
	err := r.Container.Dbr().QueryRowxContext(c, query, username).Scan(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r RepoClient) AddUser(c context.Context, m UserModel) error {
	query := "INSERT INTO accounts (username, password, email, user_role) VALUES (?, ?, ?, ?)"
	_, err := r.Container.Dbr().ExecContext(c, query, m.Username, m.Password, m.Email, m.Role)
	if err != nil {
		return err
	}
	return nil
}
