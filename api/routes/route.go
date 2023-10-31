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

	adminGroup := e.Group("/admin")
	adminGroup.Use(jwtMiddleware)
	adminGroup.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			isStudent := c.Get("is_student").(bool)
			if isStudent == true {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"message": "Not authorized as admin",
				})
			}
			return next(c)
		}
	})
	NewUserRoute(config, db, adminGroup)
	NewSubjectRoute(config, db, adminGroup)

	studentGroup := e.Group("")
	studentGroup.Use(jwtMiddleware)
	studentGroup.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			isStudent := c.Get("is_student").(bool)
			if isStudent == false {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"message": "Not authorized as student",
				})
			}
			return next(c)
		}
	})
	NewQuestionRoute(config, db, studentGroup)
}
