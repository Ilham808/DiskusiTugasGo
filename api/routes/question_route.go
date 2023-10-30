package route

import (
	"DiskusiTugas/api/controller"
	"DiskusiTugas/config"
	"DiskusiTugas/repository"
	"DiskusiTugas/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewQuestionRoute(config *config.Config, db *gorm.DB, studentGroup *echo.Group) {

	ur := repository.NewQuestionRepository(db)
	sc := controller.QuestionController{
		QuestionUseCase: usecase.NewQuestionUseCase(ur),
		Config:          config,
	}

	studentGroup.GET("/questions", sc.FetchWithPagination())
	studentGroup.GET("/questions/:id", sc.GetByID())
	studentGroup.POST("/questions", sc.Store())
	studentGroup.PUT("/questions/:id", sc.Update())
	studentGroup.DELETE("/questions/:id", sc.Destroy())
}
