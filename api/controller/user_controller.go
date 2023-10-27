package controller

import (
	"DiskusiTugas/config"
	"DiskusiTugas/domain"
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase domain.UserUseCase
	Config      *config.Config
}

type Pagination struct {
	TotalRecords int `json:"total_records"`
	CurrentPage  int `json:"current_page"`
	TotalPages   int `json:"total_pages"`
	NextPage     int `json:"next_page"`
	PrevPage     int `json:"prev_page"`
}

type UserCustomResponse struct {
	Message    string        `json:"message"`
	Data       []domain.User `json:"data"`
	Pagination Pagination    `json:"pagination"`
}

func NewUserController(userUseCase domain.UserUseCase, config *config.Config) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
		Config:      config,
	}
}

func (userController *UserController) FetchWithPagination() echo.HandlerFunc {
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

		users, totalItems, err := userController.UserUseCase.FetchWithPagination(page, pageSize)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}
		totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
		if totalPages == 0 || totalPages < page {
			totalPages = 1
		}
		prevPage := math.Max(float64(page-1), 0)

		paginationResponse := UserCustomResponse{
			Message: "Success fetch users",
			Data:    users,
			Pagination: Pagination{
				TotalRecords: totalItems,
				CurrentPage:  page,
				TotalPages:   totalPages,
				NextPage:     page + 1,
				PrevPage:     int(prevPage),
			},
		}
		return c.JSON(http.StatusOK, paginationResponse)
	}
}

func (userController *UserController) Store(user *domain.User) error {
	return userController.UserUseCase.Store(user)
}

func (userController *UserController) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid id parameter",
			})
		}

		user, err := userController.UserUseCase.GetByID(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success get user",
			"user":    user,
		})
	}

}
