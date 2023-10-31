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

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Config        *config.Config
}

func (sc *SignupController) Signup() echo.HandlerFunc {
	return func(c echo.Context) error {

		var user domain.SignupStruct

		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		validate := validator.New()
		if err := validate.Struct(user); err != nil {
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

		user.IsStudent = true

		username := user.Name
		username = strings.ReplaceAll(username, " ", "")
		baseURL := "https://robohash.org/"
		user.Avatar = baseURL + username + "?set=set4"
		user.Status = "active"

		err := sc.SignupUsecase.GetUserByEmail(user.Email)

		if err == nil {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"message": "User already exists with the given email",
			})
		}

		encryptedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(user.Password),
			bcrypt.DefaultCost,
		)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		user.Password = string(encryptedPassword)

		newUser := &domain.User{
			Name:       user.Name,
			Email:      user.Email,
			Password:   user.Password,
			Gender:     user.Gender,
			University: user.University,
			Avatar:     user.Avatar,
			IsStudent:  user.IsStudent,
			Status:     user.Status,
		}

		err = sc.SignupUsecase.Store(newUser)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		accessToken, err := sc.SignupUsecase.CreateAccessToken(newUser, sc.Config.AccessTokenSecret, sc.Config.AccessTokenExpiryHour)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		refreshToken, err := sc.SignupUsecase.CreateRefreshToken(newUser, sc.Config.RefreshTokenSecret, sc.Config.RefreshTokenExpiryHour)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Successfully register user",
			"data": map[string]interface{}{
				"access_token":  accessToken,
				"refresh_token": refreshToken,
			},
		})
	}
}
