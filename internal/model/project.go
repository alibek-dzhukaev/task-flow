package model

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name        string `gorm:"not null"`
	Description string
	Status      string `gorm:"default:active"`
	OwnerID     uint   `gorm:"not null"`

	Owner   User   `gorm:"foreignKey:OwnerID"`
	Members []User `gorm:"many2many:project_members"`
	Tasks   []Task `gorm:"foreignKey:ProjectID"`
}
