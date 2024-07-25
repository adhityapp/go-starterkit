package service

import (
	"context"
)

func (s ServiceClient) GetUserService(c context.Context) ([]UserViewModel, error) {
	var vm []UserViewModel
	m, err := s.repo.GetUser(c)
	if err != nil {
		return nil, err
	}
	for _, val := range m {
		var m UserViewModel
		m.UserID = val.UserID
		m.Password = val.Password
		m.Email = val.Email
		m.Username = val.Username
		vm = append(vm, m)
	}
	return vm, nil
}
