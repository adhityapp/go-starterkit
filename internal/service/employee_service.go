package service

import (
	"context"

	"github.com/adhityapp/go-starterkit/internal/viewmodel"
)

func (s ServiceClient) GetSmithActiveEmployee(c context.Context) ([]viewmodel.EmployeeNameViewModel, error) {
	var vm []viewmodel.EmployeeNameViewModel
	m, err := s.repo.GetSmithActiveEmployee(c)
	if err != nil {
		return nil, err
	}
	for _, val := range m {
		var m viewmodel.EmployeeNameViewModel
		m.Firstname = val.Firstname
		m.Lastname = val.Lastname
		vm = append(vm, m)
	}
	return vm, nil
}

func (s ServiceClient) GetEmployeeNoReview(c context.Context) ([]viewmodel.EmployeeNameViewModel, error) {
	var vm []viewmodel.EmployeeNameViewModel
	m, err := s.repo.GetEmployeeNoReview(c)
	if err != nil {
		return nil, err
	}
	for _, val := range m {
		var m viewmodel.EmployeeNameViewModel
		m.Firstname = val.Firstname
		m.Lastname = val.Lastname
		vm = append(vm, m)
	}
	return vm, nil
}

func (s ServiceClient) GetEmployeeDifferentDay(c context.Context) (*int, error) {
	m, err := s.repo.GetEmployeeDifferentDay(c)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (s ServiceClient) GetSalary(c context.Context) ([]viewmodel.EmployeeViewModel, error) {
	var vm []viewmodel.EmployeeViewModel
	m, err := s.repo.GetSalary(c)
	if err != nil {
		return nil, err
	}
	for _, val := range m {
		var m viewmodel.EmployeeViewModel
		m.Firstname = val.Firstname
		m.Lastname = val.Lastname
		m.EmployeeID = val.EmployeeID
		m.Salary = val.Salary
		m.SalaryNow = val.SalaryNow
		m.ReviewCount = val.ReviewCount
		vm = append(vm, m)
	}
	return vm, nil
}
