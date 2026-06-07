package repository

import (
	"github.com/alibek-dzhukaev/task-flow/internal/model"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	Create(project *model.Project) error
	FindByID(id uint) (*model.Project, error)
	FindByOwnerID(ownerID uint) ([]model.Project, error)
	Update(project *model.Project) error
	Delete(id uint) error
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) Create(project *model.Project) error {
	return r.db.Create(project).Error
}

func (r *projectRepository) FindByID(id uint) (*model.Project, error) {
	var project model.Project
	err := r.db.Preload("Owner").Preload("Members").First(&project, id).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *projectRepository) FindByOwnerID(ownerId uint) ([]model.Project, error) {
	var projects []model.Project
	err := r.db.Where("owner_id = ?", ownerId).Find(&projects).Error
	return projects, err
}

func (r *projectRepository) Update(project *model.Project) error {
	return r.db.Save(project).Error
}

func (r *projectRepository) Delete(id uint) error {
	return r.db.Delete(&model.Project{}, id).Error
}
