package repo

import (
	"context"
)

func (r RepoClient) GetUser(c context.Context) ([]UserModel, error) {
	var um []UserModel
	query := "select user_id, username, password, email from accounts"
	rows, err := r.DB.QueryContext(c, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var m UserModel
		rows.Scan(
			m.UserID,
			m.Username,
			m.Password,
			m.Email,
		)
		um = append(um, m)
	}
	return um, nil
}
