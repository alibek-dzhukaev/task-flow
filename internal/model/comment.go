package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Content  string `gorm:"not null"`
	TaskID   uint   `gorm:"not null"`
	AuthorID uint   `gorm:"not null"`

	Task   Task `gorm:"foreignKey:TaskID"`
	Author User `gorm:"foreignKey:AuthorID"`
}
