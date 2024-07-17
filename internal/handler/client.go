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
}
