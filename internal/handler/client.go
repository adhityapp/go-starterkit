package handler

import (
	"github.com/adhityapp/go-starterkit/internal/service"
	"github.com/labstack/echo"
)

type HandlerClient struct {
	service service.ServiceClient
}

func Handler(service service.ServiceClient) HandlerClient {
	return HandlerClient{
		service: service,
	}
}

func RestRegister(e *echo.Echo, h HandlerClient) {
	e.POST("/getuser", h.GetUser)
	e.POST("/getinvoice", h.GetInvoice)
	e.POST("/getinvoicedetail", h.GetInvoiceDetail)
	e.POST("/getdropdown", h.GetDropdown)
	e.POST("/deletedetailitem", h.DeleteDetailItem)
	e.POST("/addinvoice", h.AddInvoice)
	e.POST("/updateinvoice", h.UpdateInvoice)
	e.POST("/showinvoice", h.ShowInvoice)
}
