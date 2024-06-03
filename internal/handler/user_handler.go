package handler

import (
	"net/http"

	"github.com/adhityapp/go-starterkit/internal/service"
	"github.com/labstack/echo"
)

type HandlerClient struct {
	service service.ServiceClient
}

func UserHandler(service service.ServiceClient) HandlerClient {
	return HandlerClient{
		service: service,
	}
}

func RestRegister(e *echo.Echo, h HandlerClient) {
	e.POST("/getuser", h.GetUser)
}

func (uh HandlerClient) GetUser(ectx echo.Context) error {
	ctx := ectx.Request().Context()
	data, err := uh.service.GetUserService(ctx)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ectx.JSON(http.StatusOK, data)
}
