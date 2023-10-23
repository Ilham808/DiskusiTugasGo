package controller

import (
	"DiskusiTugas/config"
	"DiskusiTugas/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Config        *config.Config
}

func (sc *SignupController) Signup() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello World",
		})
	}
}
