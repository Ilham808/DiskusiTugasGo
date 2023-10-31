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
	"golang.org/x/crypto/bcrypt"
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

type UserResponsePagination struct {
	Message    string         `json:"message"`
	Data       []UserResponse `json:"data"`
	Pagination Pagination     `json:"pagination"`
}

type UserResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Status string `json:"status"`
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

		var userResponses []UserResponse
		for _, user := range users {
			userResponses = append(userResponses, UserResponse{
				ID:     user.ID,
				Name:   user.Name,
				Email:  user.Email,
				Gender: user.Gender,
				Status: user.Status,
			})
		}

		paginationResponse := UserResponsePagination{
			Message: "Success fetch users",
			Data:    userResponses,
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

func (userController *UserController) Store() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userRequest = new(domain.UserRequest)

		if err := c.Bind(&userRequest); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		validate := validator.New()
		if err := validate.Struct(userRequest); err != nil {
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

		checkEmail := userController.UserUseCase.GetByEmail(userRequest.Email)
		if checkEmail == nil {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"message": "User already exists with the given email",
			})
		}

		username := userRequest.Name
		username = strings.ReplaceAll(username, " ", "")
		baseURL := "https://robohash.org/"
		avatar := baseURL + username + "?set=set4"

		encryptedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(userRequest.Password),
			bcrypt.DefaultCost,
		)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		userRequest.Password = string(encryptedPassword)

		user := &domain.User{
			Name:       userRequest.Name,
			Email:      userRequest.Email,
			Password:   userRequest.Password,
			Gender:     userRequest.Gender,
			University: "",
			Avatar:     avatar,
			IsStudent:  false,
			Status:     userRequest.Status,
		}

		err = userController.UserUseCase.Store(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "User created successfully",
		})

	}

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
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "User not found",
			})
		}

		var userResponse UserResponse
		userResponse.Name = user.Name
		userResponse.Email = user.Email
		userResponse.Gender = user.Gender
		userResponse.Status = user.Status

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success get user",
			"user":    userResponse,
		})
	}

}

func (userController *UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userRequest = new(domain.UserRequest)

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid id parameter",
			})
		}

		if err := c.Bind(&userRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		getUserById, err := userController.UserUseCase.GetByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "User not found",
			})
		}

		validate := validator.New()
		if err := validate.Struct(userRequest); err != nil {
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

		getUserById.Name = userRequest.Name
		getUserById.Email = userRequest.Email
		getUserById.Gender = userRequest.Gender
		if userRequest.Password != "" {
			getUserById.Password = userRequest.Password
		}
		getUserById.Status = userRequest.Status

		err = userController.UserUseCase.Update(&getUserById)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "User updated successfully",
		})
	}
}

func (userController *UserController) Destroy() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid id parameter",
			})
		}

		err = userController.UserUseCase.Destroy(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "User deleted successfully",
		})
	}
}

func (userController *UserController) FecthStudent() echo.HandlerFunc {
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

		users, totalItems, err := userController.UserUseCase.FecthStudent(page, pageSize)
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

		var userResponses []UserResponse
		for _, user := range users {
			userResponses = append(userResponses, UserResponse{
				ID:     user.ID,
				Name:   user.Name,
				Email:  user.Email,
				Gender: user.Gender,
				Status: user.Status,
			})
		}

		paginationResponse := UserResponsePagination{
			Message: "Success fetch students",
			Data:    userResponses,
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

func (userController *UserController) BlockStudent() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid id parameter",
			})
		}

		err = userController.UserUseCase.BlockStudent(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Student blocked successfully",
		})
	}
}
