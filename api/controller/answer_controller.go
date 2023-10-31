package controller

import (
	"DiskusiTugas/config"
	"DiskusiTugas/domain"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AnswerController struct {
	AnswerUseCase domain.AnswerUsecase
	Config        *config.Config
}

func NewAnswerController(answerUseCase domain.AnswerUsecase, config *config.Config) *AnswerController {
	return &AnswerController{
		AnswerUseCase: answerUseCase,
		Config:        config,
	}
}

func (controller *AnswerController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		iquestionID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid id parameter",
			})
		}

		var answerRequest domain.AnswerRequest
		if err := c.Bind(&answerRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		validate := validator.New()
		if err := validate.Struct(answerRequest); err != nil {
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

		answerRequest.QuestionID = uint(iquestionID)
		xUserID, _ := strconv.ParseUint(c.Get("x-user-id").(string), 10, 32)
		answerRequest.UserID = uint(xUserID)

		file, err := c.FormFile("file")
		if err == http.ErrMissingFile {
			answerRequest.FileUrl = ""
		} else if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		} else {
			if file != nil {
				openedFile, err := file.Open()
				if err != nil {
					return c.JSON(http.StatusBadRequest, map[string]interface{}{
						"message": err.Error(),
					})
				}

				imageUrl, err := controller.AnswerUseCase.StoreFile(&domain.AnswerRequestFile{
					File: openedFile,
				})
				if err != nil {
					return c.JSON(http.StatusBadRequest, map[string]interface{}{
						"message": err.Error(),
					})
				}

				answerRequest.FileUrl = imageUrl
			}
		}

		_, err = controller.AnswerUseCase.Store(&answerRequest)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "Answer created successfully",
		})
	}
}

func (controller *AnswerController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid id parameter",
			})
		}

		var answerRequest domain.AnswerRequest
		if err := c.Bind(&answerRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		validate := validator.New()
		if err := validate.Struct(answerRequest); err != nil {
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

		xUserID, _ := strconv.ParseUint(c.Get("x-user-id").(string), 10, 32)
		answerRequest.UserID = uint(xUserID)

		file, err := c.FormFile("file")
		if err == http.ErrMissingFile {
			answerRequest.FileUrl = ""
		} else if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		} else {
			if file != nil {
				openedFile, err := file.Open()
				if err != nil {
					return c.JSON(http.StatusBadRequest, map[string]interface{}{
						"message": err.Error(),
					})
				}

				imageUrl, err := controller.AnswerUseCase.StoreFile(&domain.AnswerRequestFile{
					File: openedFile,
				})
				if err != nil {
					return c.JSON(http.StatusBadRequest, map[string]interface{}{
						"message": err.Error(),
					})
				}

				answerRequest.FileUrl = imageUrl
			}
		}

		err = controller.AnswerUseCase.Update(id, &answerRequest)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Answer updated successfully",
		})
	}
}

func (controller *AnswerController) Destroy() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid id parameter",
			})
		}

		idLogin, _ := strconv.Atoi(c.Get("x-user-id").(string))

		err = controller.AnswerUseCase.Destroy(id, uint(idLogin))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Answer deleted successfully",
		})
	}
}

func (controller *AnswerController) MarkAsCorrect() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid id parameter",
			})
		}

		idLogin, _ := strconv.Atoi(c.Get("x-user-id").(string))
		if err := controller.AnswerUseCase.MarkAsCorrect(id, uint(idLogin)); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Answer marked as correct successfully",
		})
	}
}

func (controller *AnswerController) UpVote() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid id parameter",
			})
		}

		idLogin, _ := strconv.Atoi(c.Get("x-user-id").(string))

		if err := controller.AnswerUseCase.UpVote(id, uint(idLogin)); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Answer upvoted successfully",
		})

	}
}

func (controller *AnswerController) DownVote() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid id parameter",
			})
		}

		idLogin, _ := strconv.Atoi(c.Get("x-user-id").(string))

		if err := controller.AnswerUseCase.DownVote(id, uint(idLogin)); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Answer downvoted successfully",
		})
	}
}
