package controller

import (
	"DiskusiTugas/domain"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type OpenAIController struct {
	OpenAIUseCase domain.OpenAIUseCase
}

func NewOpenAIController(openAIUseCase domain.OpenAIUseCase) *OpenAIController {
	return &OpenAIController{
		OpenAIUseCase: openAIUseCase,
	}
}

func (controller *OpenAIController) GenerateAnswer() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req domain.OpenAIRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		validate := validator.New()
		if err := validate.Struct(req); err != nil {
			var errorMsgs []string
			for _, err := range err.(validator.ValidationErrors) {
				field := strings.ToLower(err.StructNamespace())
				errorMsgs = append(errorMsgs, "Field "+field+" no valid")
			}
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Validation failed",
				"error":   errorMsgs,
			})
		}

		answer, err := controller.OpenAIUseCase.GenerateAnswer(&req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success generate answer",
			"data":    answer,
		})

	}
}
