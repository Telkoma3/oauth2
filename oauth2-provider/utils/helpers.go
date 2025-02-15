package utils

import (
    "crypto/rand"
    "encoding/base64"
    "errors"
)

// GenerateRandomString generates a random string of the given length
func GenerateRandomString(length int) (string, error) {
    b := make([]byte, length)
    _, err := rand.Read(b)
    if err != nil {
        return "", errors.New("failed to generate random string")
    }
    return base64.URLEncoding.EncodeToString(b)[:length], nil
}