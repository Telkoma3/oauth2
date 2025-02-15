package middlewares

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func RateLimitMiddleware() echo.MiddlewareFunc {
    return middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
        10, // Max 10 requests per second
    ))
}