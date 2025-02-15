package models

import "time"

type Token struct {
    ID           uint      `gorm:"primaryKey"`
    UserID       uint      `gorm:"not null"`
    AccessToken  string    `gorm:"not null"`
    RefreshToken string    `gorm:"not null"`
    ExpiresAt    time.Time `gorm:"not null"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
}