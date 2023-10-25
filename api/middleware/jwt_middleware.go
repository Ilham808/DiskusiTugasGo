package middleware

import (
	"DiskusiTugas/internal"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func JwtAuthMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			t := strings.Split(authHeader, " ")
			if len(t) == 2 {
				authToken := t[1]
				authorized, err := internal.IsAuthorized(authToken, secret)
				if err != nil {
					return c.JSON(http.StatusUnauthorized, map[string]interface{}{
						"message": err.Error(),
					})
				}
				if authorized {
					isStudent, ok := internal.ExtractIsStudentFromToken(authToken, secret)
					if !ok {
						return c.JSON(http.StatusUnauthorized, map[string]interface{}{
							"message": "Invalid Token",
						})
					}

					userID, err := internal.ExtractIDFromToken(authToken, secret)
					if err != nil {
						return c.JSON(http.StatusUnauthorized, map[string]interface{}{
							"message": err.Error(),
						})
					}
					c.Set("x-user-id", userID)
					c.Set("is_student", isStudent)
					return next(c)
				}
			}
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Not authorized",
			})
		}
	}
}
