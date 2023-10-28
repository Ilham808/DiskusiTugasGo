package route

import (
	"DiskusiTugas/api/controller"
	"DiskusiTugas/config"
	"DiskusiTugas/repository"
	"DiskusiTugas/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewSubjectRoute(config *config.Config, db *gorm.DB, adminGroup *echo.Group) {
	ur := repository.NewSubjectRepository(db)
	sc := controller.SubjectController{
		SubjectUseCase: usecase.NewSubjectUseCase(ur),
		Config:         config,
	}

	adminGroup.GET("/subjects", sc.Fetch())
	adminGroup.GET("/subjects/:id", sc.GetByID())
	adminGroup.POST("/subjects", sc.Store())
	adminGroup.PUT("/subjects/:id", sc.Update())
	adminGroup.DELETE("/subjects/:id", sc.Destroy())
}
