package service

import (
	"errors"

	"github.com/alibek-dzhukaev/task-flow/internal/model"
	"github.com/alibek-dzhukaev/task-flow/internal/repository"
	"gorm.io/gorm"
)

type ProjectService interface {
	Create(name, description string, ownerId uint) (*model.Project, error)
	GetById(id, requestID uint) (*model.Project, error)
	GetAllByOwner(ownerId uint) ([]model.Project, error)
	Update(id uint, name, description string, ownerId uint) (*model.Project, error)
	Delete(id, ownerId uint) error
}

type projectService struct {
	projectRepo repository.ProjectRepository
}

func NewProjectService(projectRepo repository.ProjectRepository) ProjectService {
	return &projectService{projectRepo: projectRepo}
}

func (s *projectService) Create(name, description string, ownerId uint) (*model.Project, error) {
	project := &model.Project{
		Name:        name,
		Description: description,
		OwnerID:     ownerId,
	}
	if err := s.projectRepo.Create(project); err != nil {
		return nil, err
	}
	return project, nil
}

func (s *projectService) GetById(id, requestID uint) (*model.Project, error) {
	project, err := s.projectRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("project not found")
		}
		return nil, err
	}
	return project, nil
}

func (s *projectService) GetAllByOwner(ownerId uint) ([]model.Project, error) {
	return s.projectRepo.FindByOwnerID(ownerId)
}

func (s *projectService) Update(id uint, name, description string, ownerId uint) (*model.Project, error) {
	project, err := s.projectRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if project.OwnerID != ownerId {
		return nil, errors.New("project not found1")
	}
	project.Name = name
	project.Description = description
	if err := s.projectRepo.Update(project); err != nil {
		return nil, err
	}
	return project, nil
}

func (s *projectService) Delete(id, ownerId uint) error {
	project, err := s.projectRepo.FindByID(id)
	if err != nil {
		return errors.New("project not found")
	}
	if project.OwnerID != ownerId {
		return errors.New("forbidden")
	}
	return s.projectRepo.Delete(id)
}
