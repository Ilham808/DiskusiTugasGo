package controller

import (
	"DiskusiTugas/config"
	"DiskusiTugas/domain"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type QuestionResponsePagination struct {
	Message    string            `json:"message"`
	Data       []domain.Question `json:"data"`
	Pagination Pagination        `json:"pagination"`
}

type QuestionController struct {
	QuestionUseCase domain.QuestionUseCase
	Config          *config.Config
}

func NewQuestionController(questionUseCase domain.QuestionUseCase, config *config.Config) *QuestionController {
	return &QuestionController{
		QuestionUseCase: questionUseCase,
		Config:          config,
	}
}

func (sc *QuestionController) FetchWithPagination() echo.HandlerFunc {
	return func(c echo.Context) error {
		pageStr := c.QueryParam("page")
		pageSizeStr := c.QueryParam("page_size")

		var page, pageSize int
		var err error

		if pageStr != "" {
			page, err = strconv.Atoi(pageStr)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"message": "Invalid page parameter",
				})
			}
		} else {
			page = 1
		}

		if pageSizeStr != "" {
			pageSize, err = strconv.Atoi(pageSizeStr)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"message": "Invalid page_size parameter",
				})
			}
		} else {
			pageSize = 10
		}

		questions, totalRecords, err := sc.QuestionUseCase.FetchWithPagination(page, pageSize)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))
		if totalPages == 0 || totalPages < page {
			totalPages = 1
		}
		prevPage := math.Max(float64(page-1), 0)

		return c.JSON(http.StatusOK, QuestionResponsePagination{
			Message: "Success get questions",
			Data:    questions,
			Pagination: Pagination{
				TotalRecords: totalRecords,
				CurrentPage:  page,
				TotalPages:   totalPages,
				NextPage:     page + 1,
				PrevPage:     int(prevPage),
			},
		})
	}
}

func (sc *QuestionController) Store() echo.HandlerFunc {
	return func(c echo.Context) error {
		var questionRequest domain.QuestionRequest
		if err := c.Bind(&questionRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		validate := validator.New()
		if err := validate.Struct(questionRequest); err != nil {
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

		file, err := c.FormFile("file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}
		if file != nil {
			openedFile, err := file.Open()
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"message": err.Error(),
				})
			}

			imageUrl, err := sc.QuestionUseCase.StoreFile(&domain.QuestionRequestFile{
				File: openedFile,
			})
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"message": err.Error(),
				})
			}

			questionRequest.FileUrl = imageUrl
		}

		userIDStr := c.Get("x-user-id").(string)
		userID, _ := strconv.ParseUint(userIDStr, 10, 32)

		questionRequest.UserID = uint(userID)

		err2 := sc.QuestionUseCase.Store(&questionRequest)
		if err2 != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Question created successfully",
		})
	}
}
