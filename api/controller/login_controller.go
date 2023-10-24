package controller

import (
	"DiskusiTugas/config"
	"DiskusiTugas/domain"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Config       *config.Config
}

func (lo *LoginController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request domain.LoginRequest
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		validate := validator.New()
		if err := validate.Struct(request); err != nil {
			var errorMsgs []string
			for _, err := range err.(validator.ValidationErrors) {
				field := strings.ToLower(err.StructNamespace())
				errorMsgs = append(errorMsgs, "Field "+field+" tidak valid")
			}
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Validation failed",
				"error":   errorMsgs,
			})
		}

		user, err := lo.LoginUsecase.GetUserByEmail(request.Email)

		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "User not found",
			})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Invalid password",
			})
		}

		accessToken, err := lo.LoginUsecase.CreateAccessToken(&user, lo.Config.AccessTokenSecret, lo.Config.AccessTokenExpiryHour)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		refreshToken, err := lo.LoginUsecase.CreateRefreshToken(&user, lo.Config.RefreshTokenSecret, lo.Config.RefreshTokenExpiryHour)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Login success",
			"data": map[string]interface{}{
				"access_token":  accessToken,
				"refresh_token": refreshToken,
			},
		})

	}
}
