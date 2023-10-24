package route

import (
	"DiskusiTugas/config"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoute(e *echo.Echo, config *config.Config, db *gorm.DB) {
	NewSignupRouter(e, config, db)
}
