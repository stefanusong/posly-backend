package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	role      Role   `gorm:"not null"`
	firstName string `gorm:"not null"`
	lastName  string
	email     string `gorm:"unique;not null"`
	gender    string
	dob       time.Time `gorm:"not null"`
	password  string    `gorm:"not null"`
}
