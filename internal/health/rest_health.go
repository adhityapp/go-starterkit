package health

import (
	"net/http"

	"github.com/adhityapp/go-starterkit/bootstrap"
	"github.com/adhityapp/go-starterkit/rest/health"
	"github.com/labstack/echo"
)

type RestService struct{}

func RestRegister(cfg *bootstrap.Container, e *echo.Echo) {
	health.Router(e, &RestService{})
}

func (s *RestService) GetHealth(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, ".")
}

func (s *RestService) GetHealthMixin0(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, ".")
}
