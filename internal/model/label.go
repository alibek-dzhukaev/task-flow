package model

import "time"

type Label struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name      string `gorm:"not null"`
	Color     string `gorm:"default:#6366f1"`
	ProjectID uint   `gorm:"not null"`

	Project Project `gorm:"foreignKey:ProjectID"`
}
