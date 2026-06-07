package middlewares

import (
	"backendmaw/dto"
	"errors"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		tokenCookie, err := c.Cookie("token")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, new(dto.FailedResponse("Must Login!", http.StatusUnauthorized)))
		}
		jwtKey := []byte(os.Getenv("JWT_KEY"))
		token, err := jwt.ParseWithClaims(tokenCookie.Value, &dto.UserToken{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Unexpected signing method: " + token.Header["alg"].(string))
			}
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, new(dto.FailedResponse("Token Invalid!", http.StatusUnauthorized)))
		}
		claims, _ := token.Claims.(*dto.UserToken)
		c.Set("user", claims)
		return next(c)
	}
}
