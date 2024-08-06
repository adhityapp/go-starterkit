package service

import (
	"context"

	"github.com/adhityapp/go-starterkit/internal/repo"
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

func (s ServiceClient) LoginService(c context.Context, username string) (*viewmodel.UserViewModel, error) {
	m, err := s.repo.GetUserbyUsername(c, username)
	if err != nil {
		return nil, err
	}
	vm := viewmodel.UserViewModel{
		UserID:   m.UserID,
		Username: m.Username,
		Password: m.Password,
		Email:    m.Email,
		Role:     m.Role,
	}

	return &vm, nil
}

func (s ServiceClient) AddUserService(c context.Context, vm viewmodel.UserViewModel) error {
	m := repo.UserModel{
		Username: vm.Username,
		Password: vm.Password,
		Email:    vm.Email,
		Role:     vm.Role,
	}

	err := s.repo.AddUser(c, m)
	if err != nil {
		return err
	}

	return nil
}
