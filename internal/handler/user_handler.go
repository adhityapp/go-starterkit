package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func (uh HandlerClient) GetUser(ectx echo.Context) error {
	ctx := ectx.Request().Context()
	data, err := uh.service.GetUserService(ctx)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ectx.JSON(http.StatusOK, data)
}
