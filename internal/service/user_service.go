package service

import (
	"context"

	"github.com/adhityapp/go-starterkit/internal/viewmodel"
)

func (s ServiceClient) GetUserService(c context.Context) ([]viewmodel.UserViewModel, error) {
	var vm []viewmodel.UserViewModel
	m, err := s.repo.GetUser(c)
	if err != nil {
		return nil, err
	}
	for _, val := range m {
		var m viewmodel.UserViewModel
		m.UserID = val.UserID
		m.Password = val.Password
		m.Email = val.Email
		m.Username = val.Username
		vm = append(vm, m)
	}
	return vm, nil
}
