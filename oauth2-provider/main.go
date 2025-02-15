package main

import (
    "oauth2-provider/config"
    "oauth2-provider/routes"
    "oauth2-provider/utils"
    "github.com/labstack/echo/v4"
)

func main() {
    utils.InitLogger()

    config.InitDB()
    config.InitRedis()

    e := echo.New()

    // Set custom error handler
    e.HTTPErrorHandler = utils.ErrorHandler

    routes.SetupRoutes(e)

    services.InitOAuth2Service(config.RedisClient)

    utils.Logger().Info("Starting OAuth2 Provider on :8080")
    e.Start(":8080")
}