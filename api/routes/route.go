package route

import (
	"DiskusiTugas/api/middleware"
	"DiskusiTugas/config"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoute(e *echo.Echo, config *config.Config, db *gorm.DB) {
	NewSignupRouter(e, config, db)
	NewLoginRouter(e, config, db)

	jwtMiddleware := middleware.JwtAuthMiddleware(config.AccessTokenSecret)

	adminGroup := e.Group("")
	adminGroup.Use(jwtMiddleware)
	adminGroup.GET("/hello", func(c echo.Context) error {
		isStudent := c.Get("is_student").(bool)
		if isStudent {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Not authorized as admin",
			})
		}
		return c.String(http.StatusOK, "Hello, World!")
	})

	studentGroup := e.Group("/student")
	studentGroup.Use(jwtMiddleware)
	studentGroup.GET("/hello", func(c echo.Context) error {
		isStudent := c.Get("is_student").(bool)
		if !isStudent {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Not authorized as student",
			})
		}
		return c.String(http.StatusOK, "Hello, World!")
	})
}
