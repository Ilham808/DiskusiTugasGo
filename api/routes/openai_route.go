package route

import (
	"DiskusiTugas/api/controller"
	"DiskusiTugas/config"
	"DiskusiTugas/repository"
	"DiskusiTugas/usecase"

	"github.com/labstack/echo/v4"
)

func NewOpenAiRoute(config *config.Config, studentGroup *echo.Group) {
	ur := repository.NewOpenAIRepository()
	sc := controller.NewOpenAIController(
		usecase.NewOpenAIUseCase(ur),
	)
	studentGroup.POST("/tanya-bot", sc.GenerateAnswer())
}
