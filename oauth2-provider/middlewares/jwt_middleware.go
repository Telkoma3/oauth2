package middlewares

import (
    "net/http"

    "github.com/dgrijalva/jwt-go"
    "github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        tokenString := c.Request().Header.Get("Authorization")
        if tokenString == "" {
            return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing token"})
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte("your-secret-key"), nil
        })

        if err != nil || !token.Valid {
            return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
        }

        claims := token.Claims.(jwt.MapClaims)
        c.Set("user_id", claims["user_id"])

        return next(c)
    }
}