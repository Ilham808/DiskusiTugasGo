package route

import (
	"DiskusiTugas/api/controller"
	"DiskusiTugas/config"
	"DiskusiTugas/repository"
	"DiskusiTugas/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewUserRoute(config *config.Config, db *gorm.DB, adminGroup *echo.Group) {
	ur := repository.NewUserRepository(db)
	sc := controller.UserController{
		UserUseCase: usecase.NewUserUseCase(ur),
		Config:      config,
	}

	adminGroup.GET("/users", sc.FetchWithPagination())
	adminGroup.GET("/users/:id", sc.GetByID())
	adminGroup.POST("/users", sc.Store())
	adminGroup.PUT("/users/:id", sc.Update())
	adminGroup.DELETE("/users/:id", sc.Destroy())
}
