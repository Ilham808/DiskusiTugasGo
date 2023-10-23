package route

import (
	"DiskusiTugas/api/controller"
	"DiskusiTugas/config"
	"DiskusiTugas/repository"
	"DiskusiTugas/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewSignupRouter(e *echo.Echo, config *config.Config, db *gorm.DB) {
	ur := repository.NewUserRepository(db)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur),
		Config:        config,
	}
	e.POST("/signup", sc.Signup())
}
