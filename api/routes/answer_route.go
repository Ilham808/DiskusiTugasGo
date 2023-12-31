package route

import (
	"DiskusiTugas/api/controller"
	"DiskusiTugas/config"
	"DiskusiTugas/repository"
	"DiskusiTugas/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewAnswerRoute(config *config.Config, db *gorm.DB, studentGroup *echo.Group) {
	ur := repository.NewAnswerRepository(db)
	sc := controller.AnswerController{
		AnswerUseCase: usecase.NewAnswerUseCase(ur),
		Config:        config,
	}

	studentGroup.POST("/questions/:id/answers", sc.Create())
	studentGroup.PUT("/answers/:id", sc.Update())
	studentGroup.DELETE("/answers/:id", sc.Destroy())
	studentGroup.PUT("/answers/:id/correct", sc.MarkAsCorrect())

	studentGroup.PUT("/answers/:id/upvote", sc.UpVote())
	studentGroup.PUT("/answers/:id/downvote", sc.DownVote())

	studentGroup.POST("/answers/:id/comments", sc.Comment())
	studentGroup.DELETE("/answers/:id/comments/:comment_id", sc.DestroyComment())
}
