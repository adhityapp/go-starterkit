package service

import (
	"context"
	"time"

	"github.com/adhityapp/go-starterkit/internal/repo"
	"github.com/adhityapp/go-starterkit/internal/viewmodel"
	"github.com/adhityapp/go-starterkit/pkg/utils"
)

func (s ServiceClient) GetinvoiceService(c context.Context) ([]InvoiceViewModel, error) {
	var vm []InvoiceViewModel
	m, err := s.repo.GetInvoice(c)
	if err != nil {
		utils.LogError("service", "get-repo", err)
		return nil, err
	}
	for _, val := range m {
		var m InvoiceViewModel
		m.InvoiceID = val.InvoiceID
		m.IssueDate, _ = time.Parse("2006-01-02 15:04:05", string(val.IssueDate))
		m.Subject = val.Subject
		m.TotalItems = val.TotalItems
		m.CustomerID = val.CustomerID
		m.Amount = val.Amount
		m.DueDate, _ = time.Parse("2006-01-02 15:04:05", string(val.DueDate))
		m.Status = val.Status
		vm = append(vm, m)
	}
	return vm, nil
}

func (s ServiceClient) GetinvoiceDetailService(c context.Context, id int) (*InvoiceDetailViewModel, error) {
	var vm []DetailItemViewModel2
	m, err := s.repo.GetInvoiceDetail(c, id)
	if err != nil {
		utils.LogError("service", "get-repo", err)
		return nil, err
	}
	for _, val := range m {
		var m DetailItemViewModel2
		m.DetailItemID = val.DetailItemID
		m.ItemID = val.ItemID
		m.Qty = val.Qty
		m.Amount = val.Amount
		vm = append(vm, m)
	}

	inv, err := s.repo.GetInvoiceByID(c, id)
	if err != nil {
		return nil, err
	}

	isdt, _ := time.Parse("2006-01-02 15:04:05", string(inv.IssueDate))
	ddt, _ := time.Parse("2006-01-02 15:04:05", string(inv.DueDate))

	idvm := InvoiceDetailViewModel{
		InvoiceID:  inv.InvoiceID,
		IssueDate:  isdt,
		Subject:    inv.Subject,
		TotalItems: inv.TotalItems,
		CustomerID: inv.CustomerID,
		Amount:     inv.Amount,
		DueDate:    ddt,
		Status:     inv.Status,
		Detail:     vm,
	}
	return &idvm, nil
}

func (s ServiceClient) GetDropdownService(c context.Context) (*DropdownViewModel, error) {
	var cvm []CustomerViewModel
	m, err := s.repo.GetCustomer(c)
	if err != nil {
		utils.LogError("service", "get-repo", err)
		return nil, err
	}
	for _, val := range m {
		var m CustomerViewModel
		m.CustomerID = val.CustomerID
		m.CustomerName = val.CustomerName
		m.CustomerAddress = val.CustomerAddress
		cvm = append(cvm, m)
	}

	var ivm []ItemViewModel
	i, err := s.repo.GetItem(c)
	if err != nil {
		utils.LogError("service", "get-repo", err)
		return nil, err
	}
	for _, val := range i {
		var m ItemViewModel
		m.ItemID = val.ItemID
		m.ItemName = val.ItemName
		m.ItemType = val.ItemType
		m.UnitPrice = val.UnitPrice
		ivm = append(ivm, m)
	}

	dd := DropdownViewModel{
		Customer: cvm,
		Items:    ivm,
	}
	return &dd, nil
}

func (s ServiceClient) DeleteDetailItemService(c context.Context, id int) error {
	err := s.repo.DeleteDetailItem(c, id)
	if err != nil {
		utils.LogError("service", "get-repo", err)
		return err
	}
	return nil
}

func (s ServiceClient) AddInvoiceService(c context.Context, vm viewmodel.RequestInvoice) error {
	im := repo.InvoiceModelInsert{
		IssueDate:  vm.IssueDate,
		Subject:    vm.Subject,
		TotalItems: vm.TotalItems,
		CustomerID: vm.CustomerID,
		Amount:     vm.Amount,
		DueDate:    vm.DueDate,
	}
	invid, err := s.repo.AddInvoice(c, im)
	if err != nil {
		utils.LogError("service", "get-repo", err)
		return err
	}
	var dm []repo.DetailItemModel
	for _, val := range vm.Detail {
		m := repo.DetailItemModel{
			InvoiceID: *invid,
			ItemID:    val.ItemID,
			Qty:       val.Qty,
			Amount:    val.Amount,
		}
		dm = append(dm, m)
	}
	err = s.repo.AddInvoiceDetail(c, dm)
	if err != nil {
		return err
	}
	return nil
}

func (s ServiceClient) UpdateInvoiceService(c context.Context, vm viewmodel.RequestInvoice) error {
	im := repo.InvoiceModelInsert{
		InvoiceID:  vm.InvoiceID,
		IssueDate:  vm.IssueDate,
		Subject:    vm.Subject,
		TotalItems: vm.TotalItems,
		CustomerID: vm.CustomerID,
		Amount:     vm.Amount,
		DueDate:    vm.DueDate,
	}
	err := s.repo.UpdateInvoice(c, im)
	if err != nil {
		utils.LogError("service", "get-repo", err)
		return err
	}
	var dm []repo.DetailItemModel
	for _, val := range vm.Detail {
		m := repo.DetailItemModel{
			DetailItemID: val.DetailItemID,
			ItemID:       val.ItemID,
			Qty:          val.Qty,
			Amount:       val.Amount,
		}
		dm = append(dm, m)
	}
	err = s.repo.UpdateInvoiceItem(c, dm)
	if err != nil {
		return err
	}
	return nil
}

func (s ServiceClient) ShowInvoiceService(c context.Context, id int) (*InvoiceDetailViewModel, error) {
	var vm []DetailItemViewModel2
	m, err := s.repo.GetInvoiceDetail(c, id)
	if err != nil {
		utils.LogError("service", "get-repo", err)
		return nil, err
	}
	for _, val := range m {
		var m DetailItemViewModel2
		m.ItemName = val.ItemName
		m.DetailItemID = val.DetailItemID
		m.ItemID = val.ItemID
		m.Qty = val.Qty
		m.Amount = val.Amount
		vm = append(vm, m)
	}

	inv, err := s.repo.GetInvoiceByID(c, id)
	if err != nil {
		return nil, err
	}

	isdt, _ := time.Parse("2006-01-02 15:04:05", string(inv.IssueDate))
	ddt, _ := time.Parse("2006-01-02 15:04:05", string(inv.DueDate))

	idvm := InvoiceDetailViewModel{
		InvoiceID:    inv.InvoiceID,
		IssueDate:    isdt,
		Subject:      inv.Subject,
		TotalItems:   inv.TotalItems,
		CustomerID:   inv.CustomerID,
		CustomerName: inv.CustomerName,
		CustomerAddr: inv.CustomerAddress,
		Amount:       inv.Amount,
		DueDate:      ddt,
		Status:       inv.Status,
		Detail:       vm,
	}
	return &idvm, nil
}
