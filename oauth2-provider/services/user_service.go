package services

import (
    "errors"
    "golang.org/x/crypto/bcrypt"
    "oauth2-provider/config"
    "oauth2-provider/models"
)

func RegisterUser(username, password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    user := models.User{
        Username: username,
        Password: string(hashedPassword),
    }

    result := config.DB.Create(&user)
    if result.Error != nil {
        return result.Error
    }

    return nil
}

func AuthenticateUser(username, password string) (uint, error) {
    var user models.User
    result := config.DB.Where("username = ?", username).First(&user)
    if result.Error != nil {
        return 0, errors.New("user not found")
    }

    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return 0, errors.New("invalid password")
    }

    return user.ID, nil
}