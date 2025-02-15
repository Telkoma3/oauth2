package routes

import (
    "oauth2-provider/controllers"
    "github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
    e.Use(middlewares.RateLimitMiddleware())
    e.Use(middlewares.MetricsMiddleware)

    e.POST("/register", controllers.RegisterUser)
    e.POST("/login", controllers.LoginUser)
    e.GET("/authorize", controllers.Authorize)
    e.POST("/token", controllers.GetToken)
    e.GET("/protected", controllers.ProtectedEndpoint, middlewares.JWTMiddleware)
    e.GET("/users/:id", controllers.GetUser)

    middlewares.RegisterMetrics(e)
}