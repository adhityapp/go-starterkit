package handler

import (
	"net/http"

	"github.com/adhityapp/go-starterkit/internal/viewmodel"
	"github.com/adhityapp/go-starterkit/pkg/utils"
	"github.com/labstack/echo"
)

func (uh HandlerClient) GetInvoice(ectx echo.Context) error {
	ctx := ectx.Request().Context()
	data, err := uh.service.GetinvoiceService(ctx)
	if err != nil {
		utils.LogError("handler", "get-service", err)
		return ectx.JSON(http.StatusInternalServerError, err.Error())
	}
	response := Response{
		Success: true,
		Message: "Invoice Data",
		Data:    data,
	}
	return ectx.JSON(http.StatusOK, response)
}

func (uh HandlerClient) GetInvoiceDetail(ectx echo.Context) error {
	ctx := ectx.Request().Context()
	var req RequestID
	err := ectx.Bind(&req)
	if err != nil {
		utils.LogError("handler", "get-service", err)
		return ectx.JSON(http.StatusInternalServerError, err.Error())
	}
	data, err := uh.service.GetinvoiceDetailService(ctx, req.ID)
	if err != nil {
		utils.LogError("handler", "get-service", err)
		return ectx.JSON(http.StatusInternalServerError, err.Error())
	}
	response := Response{
		Success: true,
		Message: "Invoice Detail",
		Data:    data,
	}
	return ectx.JSON(http.StatusOK, response)
}

func (uh HandlerClient) GetDropdown(ectx echo.Context) error {
	ctx := ectx.Request().Context()

	data, err := uh.service.GetDropdownService(ctx)
	if err != nil {
		utils.LogError("handler", "get-service", err)
		return ectx.JSON(http.StatusInternalServerError, err.Error())
	}
	response := Response{
		Success: true,
		Message: "Dropdown Detail",
		Data:    data,
	}
	return ectx.JSON(http.StatusOK, response)
}

func (uh HandlerClient) DeleteDetailItem(ectx echo.Context) error {
	ctx := ectx.Request().Context()
	var req RequestID
	err := ectx.Bind(&req)
	if err != nil {
		utils.LogError("handler", "get-service", err)
		return ectx.JSON(http.StatusInternalServerError, err.Error())
	}
	err = uh.service.DeleteDetailItemService(ctx, req.ID)
	if err != nil {
		utils.LogError("handler", "get-service", err)
		return ectx.JSON(http.StatusInternalServerError, err.Error())
	}
	response := Response{
		Success: true,
		Message: "Item Deleted",
	}
	return ectx.JSON(http.StatusOK, response)
}

func (uh HandlerClient) AddInvoice(ectx echo.Context) error {
	ctx := ectx.Request().Context()
	var req viewmodel.RequestInvoice
	err := ectx.Bind(&req)
	if err != nil {
		utils.LogError("handler", "get-service", err)
		return ectx.JSON(http.StatusInternalServerError, err.Error())
	}
	err = uh.service.AddInvoiceService(ctx, req)
	if err != nil {
		utils.LogError("handler", "get-service", err)
		return ectx.JSON(http.StatusInternalServerError, err.Error())
	}
	response := Response{
		Success: true,
		Message: "Invoice Created",
	}
	return ectx.JSON(http.StatusOK, response)
}

func (uh HandlerClient) UpdateInvoice(ectx echo.Context) error {
	ctx := ectx.Request().Context()
	var req viewmodel.RequestInvoice
	err := ectx.Bind(&req)
	if err != nil {
		utils.LogError("handler", "get-service", err)
		return ectx.JSON(http.StatusInternalServerError, err.Error())
	}
	err = uh.service.UpdateInvoiceService(ctx, req)
	if err != nil {
		utils.LogError("handler", "get-service", err)
		return ectx.JSON(http.StatusInternalServerError, err.Error())
	}
	response := Response{
		Success: true,
		Message: "Invoice Updated",
	}
	return ectx.JSON(http.StatusOK, response)
}

func (uh HandlerClient) ShowInvoice(ectx echo.Context) error {
	ctx := ectx.Request().Context()
	var req RequestID
	err := ectx.Bind(&req)
	if err != nil {
		utils.LogError("handler", "get-service", err)
		return ectx.JSON(http.StatusInternalServerError, err.Error())
	}
	data, err := uh.service.ShowInvoiceService(ctx, req.ID)
	if err != nil {
		utils.LogError("handler", "get-service", err)
		return ectx.JSON(http.StatusInternalServerError, err.Error())
	}
	response := Response{
		Success: true,
		Message: "Invoice Detail",
		Data:    data,
	}
	return ectx.JSON(http.StatusOK, response)
}
