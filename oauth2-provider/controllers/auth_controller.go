package controllers

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "gopkg.in/oauth2.v4/server"
    "oauth2-provider/services"
)

func RegisterUser(c echo.Context) error {
    type RegisterRequest struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    var req RegisterRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
    }

    err := services.RegisterUser(req.Username, req.Password)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusCreated, map[string]string{"message": "User registered successfully"})
}

func LoginUser(c echo.Context) error {
    type LoginRequest struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    var req LoginRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
    }

    userID, err := services.AuthenticateUser(req.Username, req.Password)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "user_id": userID,
        "message": "Login successful",
    })
}

func Authorize(c echo.Context) error {
    err := services.OAuth2Server.HandleAuthorizeRequest(c.Response(), c.Request())
    if err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }
    return nil
}

func GetToken(c echo.Context) error {
    err := services.OAuth2Server.HandleTokenRequest(c.Response(), c.Request())
    if err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }
    return nil
}