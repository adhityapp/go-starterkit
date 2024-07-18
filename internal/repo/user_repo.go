package repo

import (
	"context"
)

func (r RepoClient) GetUser(c context.Context) ([]UserModel, error) {
	var um []UserModel
	query := "select user_id, username, password, email from accounts"
	rows, err := r.DB.Dbr().QueryxContext(c, query)
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
