package controller

import (
	"DiskusiTugas/config"
	"DiskusiTugas/domain"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type QuestionController struct {
	QuestionUseCase domain.QuestionUseCase
	Config          *config.Config
}

type QuestionRespon struct {
	ID          uint   `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
	UserID      uint   `json:"user_id"`
	UserName    string `json:"user_name"`
	SubjectID   uint   `json:"subject_id"`
	SubjectName string `json:"subject_name"`
	Question    string `json:"question"`
	Description string `json:"description"`
	File        string `json:"file"`
}

type QuestionResponsePagination struct {
	Message    string           `json:"message"`
	Data       []QuestionRespon `json:"data"`
	Pagination Pagination       `json:"pagination"`
}

type QuestionResponDetail struct {
	QuestionID uint   `json:"question_id"`
	Text       string `json:"text"`
	ImageURL   string `json:"image_url"`
	AskedBy    struct {
		UserID uint   `json:"user_id"`
		Name   string `json:"name"`
	} `json:"asked_by"`
	Answers []struct {
		AnswerID   uint   `json:"answer_id"`
		Text       string `json:"text"`
		ImageURL   string `json:"image_url"`
		AnsweredBy struct {
			UserID   uint   `json:"user_id"`
			UserName string `json:"user_name"`
		} `json:"answered_by"`
		IsCorrect bool `json:"is_correct"`
		Comments  []struct {
			CommentID   uint   `json:"comment_id"`
			Text        string `json:"text"`
			CommentedBy struct {
				UserID   uint   `json:"user_id"`
				UserName string `json:"user_name"`
			} `json:"commented_by"`
		} `json:"comments"`
		Votes int `json:"votes"`
	} `json:"answers"`
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

		questionResponses := []QuestionRespon{}

		for _, question := range questions {
			questionResponses = append(questionResponses, QuestionRespon{
				ID:          question.ID,
				CreatedAt:   question.CreatedAt.String(),
				UpdatedAt:   question.CreatedAt.String(),
				DeletedAt:   question.CreatedAt.String(),
				UserID:      question.UserID,
				UserName:    question.User.Name,
				SubjectID:   question.SubjectID,
				SubjectName: question.Subject.Name,
				Question:    question.Question,
				Description: question.Description,
				File:        question.File,
			})
		}

		return c.JSON(http.StatusOK, QuestionResponsePagination{
			Message: "Success get questions",
			Data:    questionResponses,
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

func (sc *QuestionController) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid id parameter",
			})
		}

		question, err := sc.QuestionUseCase.GetByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Question not found",
			})
		}

		response := QuestionResponDetail{
			QuestionID: question.ID,
			Text:       question.Question,
			ImageURL:   question.File,
			AskedBy: struct {
				UserID uint   `json:"user_id"`
				Name   string `json:"name"`
			}{
				UserID: question.UserID,
				Name:   question.User.Name,
			},
			Answers: make([]struct {
				AnswerID   uint   `json:"answer_id"`
				Text       string `json:"text"`
				ImageURL   string `json:"image_url"`
				AnsweredBy struct {
					UserID   uint   `json:"user_id"`
					UserName string `json:"user_name"`
				} `json:"answered_by"`
				IsCorrect bool `json:"is_correct"`
				Comments  []struct {
					CommentID   uint   `json:"comment_id"`
					Text        string `json:"text"`
					CommentedBy struct {
						UserID   uint   `json:"user_id"`
						UserName string `json:"user_name"`
					} `json:"commented_by"`
				} `json:"comments"`
				Votes int `json:"votes"`
			}, len(question.Answer)),
		}

		sort.Slice(response.Answers, func(i, j int) bool {
			return response.Answers[i].Votes > response.Answers[j].Votes
		})

		for i, answer := range question.Answer {
			response.Answers[i].AnswerID = answer.ID
			response.Answers[i].Text = answer.Answer
			response.Answers[i].ImageURL = answer.File
			response.Answers[i].AnsweredBy.UserID = answer.UserID
			response.Answers[i].AnsweredBy.UserName = answer.User.Name
			response.Answers[i].IsCorrect = answer.IsCorrect
			response.Answers[i].Votes = answer.Vote

			for _, comment := range answer.Comment {
				response.Answers[i].Comments = append(response.Answers[i].Comments, struct {
					CommentID   uint   `json:"comment_id"`
					Text        string `json:"text"`
					CommentedBy struct {
						UserID   uint   `json:"user_id"`
						UserName string `json:"user_name"`
					} `json:"commented_by"`
				}{
					CommentID: comment.ID,
					Text:      comment.Comment,
					CommentedBy: struct {
						UserID   uint   `json:"user_id"`
						UserName string `json:"user_name"`
					}{
						UserID:   comment.UserID,
						UserName: comment.User.Name,
					},
				})
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "Success get question",
			"question": response,
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
		if err == http.ErrMissingFile {
			questionRequest.FileUrl = ""
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

func (sc *QuestionController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var questionRequest domain.QuestionRequest

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid id parameter",
			})
		}

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

		question, err := sc.QuestionUseCase.GetByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Question not found",
			})
		}

		xUserID, _ := strconv.ParseUint(c.Get("x-user-id").(string), 10, 32)
		if question.UserID != uint(xUserID) {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "You are not authorized to update this question",
			})
		}

		file, err := c.FormFile("file")
		if err == http.ErrMissingFile {
			questionRequest.FileUrl = ""
		} else if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		} else {
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

		err2 := sc.QuestionUseCase.Update(id, &questionRequest)
		if err2 != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Question updated successfully",
		})

	}
}

func (sc *QuestionController) Destroy() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid id parameter",
			})
		}

		question, err := sc.QuestionUseCase.GetByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Question not found",
			})
		}

		xUserID, _ := strconv.ParseUint(c.Get("x-user-id").(string), 10, 32)
		if question.UserID != uint(xUserID) {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "You are not authorized to delete this question",
			})
		}

		err2 := sc.QuestionUseCase.Destroy(id)
		if err2 != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Question deleted successfully",
		})
	}
}
