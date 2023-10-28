package controller

import (
	"DiskusiTugas/config"
	"DiskusiTugas/domain"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type SubjectController struct {
	SubjectUseCase domain.SubjectUseCase
	Config         *config.Config
}

func NewSubjectController(subjectUseCase domain.SubjectUseCase, config *config.Config) *SubjectController {
	return &SubjectController{
		SubjectUseCase: subjectUseCase,
		Config:         config,
	}
}

func generateSlug(name string) string {
	numberRand := func() string {
		return strconv.Itoa(rand.Intn(1000))
	}
	name = strings.ReplaceAll(name, " ", "-")
	return fmt.Sprintf("%s-%s", strings.ToLower(name), numberRand())
}

func (subjectController *SubjectController) Fetch() echo.HandlerFunc {
	return func(c echo.Context) error {

		subjects, err := subjectController.SubjectUseCase.Fetch()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "Success get subjects",
			"subjects": subjects,
		})
	}
}

func (subjectController *SubjectController) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid id parameter",
			})
		}

		subject, err := subjectController.SubjectUseCase.GetByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Subject not found",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success get subject",
			"subject": subject,
		})
	}
}

func (subjectController *SubjectController) Store() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request domain.SubjectRequest
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
				errorMsgs = append(errorMsgs, "Field "+field+" no valid")
			}
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Validation failed",
				"error":   errorMsgs,
			})
		}

		slug := generateSlug(request.Name)

		subject := domain.Subject{
			Name: request.Name,
			Slug: slug,
		}

		if err := subjectController.SubjectUseCase.Store(&subject); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Subject created successfully",
		})
	}
}

func (subjectController *SubjectController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid id parameter",
			})
		}

		subject, err := subjectController.SubjectUseCase.GetByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Subject not found",
			})
		}

		var request domain.SubjectRequest
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
				errorMsgs = append(errorMsgs, "Field "+field+" no valid")
			}
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Validation failed",
				"error":   errorMsgs,
			})
		}

		subject.Name = request.Name
		subject.Slug = generateSlug(request.Name)

		if err := subjectController.SubjectUseCase.Update(&subject); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Subject updated successfully",
		})
	}
}

func (subjectController *SubjectController) Destroy() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid id parameter",
			})
		}

		err = subjectController.SubjectUseCase.Destroy(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Subject deleted successfully",
		})
	}
}
