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
	e.POST("/q2", h.GetSmithActiveEmployee)
	e.POST("/q3", h.GetEmployeeNoReview)
	e.POST("/q4", h.GetEmployeeDifferentDay)
	e.POST("/q5", h.GetSalary)
	e.POST("/q7", h.GetTxt)
	e.POST("/q8", h.GetCity)
	e.POST("/q9", h.GetArray)
	e.POST("/q10", h.GetRandomString)
}
