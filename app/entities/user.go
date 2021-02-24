package entities

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email           string
	Username        string
	Password        string
	RememberToken   string
	EmailVerifiedAt time.Time
}
