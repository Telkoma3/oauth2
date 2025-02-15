package utils

import (
    "fmt"
    "net/http"

    "github.com/labstack/echo/v4"
)

type CustomError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

func (e *CustomError) Error() string {
    return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func NewError(code int, message string) *CustomError {
    return &CustomError{
        Code:    code,
        Message: message,
    }
}

func ErrorHandler(err error, c echo.Context) {
    var customErr *CustomError
    if errors.As(err, &customErr) {
        c.JSON(customErr.Code, map[string]string{"error": customErr.Message})
    } else {
        c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
    }
}