package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Email     string `gorm:"uniqeIndex;not null"`
	Password  string `gorm:"not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Role      string `gorm:"default:member"`
	IsActive  bool   `gorm:"default:true"`
}
