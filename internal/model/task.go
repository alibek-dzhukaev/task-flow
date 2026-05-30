package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Title       string `gorm:"not null"`
	Description string
	Status      string `gorm:"default:todo"`
	Priority    string `gorm:"default:medium"`
	DueDate     *time.Time
	ProjectID   uint `gorm:"not null"`
	AssigneeID  *uint

	Project  Project `gorm:"foreignKey:ProjectID"`
	Assignee User    `gorm:"foreignKey:AssigneeID"`
	Labels   []Label `gorm:"many2many:task_labels"`
}
