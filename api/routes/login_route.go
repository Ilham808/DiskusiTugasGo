package route

import (
	"DiskusiTugas/api/controller"
	"DiskusiTugas/config"
	"DiskusiTugas/repository"
	"DiskusiTugas/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewLoginRouter(e *echo.Echo, config *config.Config, db *gorm.DB) {
	ur := repository.NewUserRepository(db)
	sc := controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur),
		Config:       config,
	}
	e.POST("/login", sc.Login())
}
