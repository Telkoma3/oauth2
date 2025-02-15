package controllers

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "oauth2-provider/config"
    "oauth2-provider/models"
)

func GetUser(c echo.Context) error {
    userID := c.Param("id")

    var user models.User
    result := config.DB.First(&user, userID)
    if result.Error != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
    }

    return c.JSON(http.StatusOK, user)
}