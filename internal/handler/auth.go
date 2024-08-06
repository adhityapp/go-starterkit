package handler

import (
	"net/http"

	"github.com/adhityapp/go-starterkit/auth"
	"github.com/adhityapp/go-starterkit/internal/viewmodel"
	"github.com/labstack/echo"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (uh HandlerClient) Login(ectx echo.Context) error {
	var user UserLogin
	ctx := ectx.Request().Context()

	err := ectx.Bind(&user)
	if err != nil {
		return err
	}

	data, err := uh.service.LoginService(ctx, user.Username)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, err.Error())
	}

	if user.Username != data.Username || user.Password != data.Password {
		return ectx.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid username or password"})
	}

	token, err := auth.GenerateJWT(data.Username, data.Role)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to generate token"})
	}

	return ectx.JSON(http.StatusOK, echo.Map{"token": token})
}

func (uh HandlerClient) Register(ectx echo.Context) error {
	ctx := ectx.Request().Context()

	var req viewmodel.UserViewModel
	err := ectx.Bind(&req)
	if err != nil {
		return err
	}

	err = uh.service.AddUserService(ctx, req)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ectx.JSON(http.StatusOK, echo.Map{"message": "Success Add Data"})
}
