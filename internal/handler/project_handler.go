package handler

import (
	"encoding/json"
	"net/http"

	"github.com/alibek-dzhukaev/task-flow/internal/dto"
	"github.com/alibek-dzhukaev/task-flow/internal/middleware"
	"github.com/alibek-dzhukaev/task-flow/internal/service"
	"github.com/alibek-dzhukaev/task-flow/internal/util"
)

type ProjectHandler struct {
	projectService service.ProjectService
}

func NewProjectHandler(projectService service.ProjectService) *ProjectHandler {
	return &ProjectHandler{projectService: projectService}
}

func (h *ProjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(middleware.UserClaimsKey).(*util.Claims)

	var req dto.CreateProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Fail(w, http.StatusBadRequest, "invalid request body")
		return
	}

	project, err := h.projectService.Create(req.Name, req.Description, claims.UserID)
	if err != nil {
		Fail(w, http.StatusInternalServerError, err.Error())
		return
	}
	Success(w, http.StatusCreated, project)
}

func (h *ProjectHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(middleware.UserClaimsKey).(*util.Claims)

	projects, err := h.projectService.GetAllByOwner(claims.UserID)
	if err != nil {
		Fail(w, http.StatusInternalServerError, err.Error())
		return
	}
	Success(w, http.StatusOK, projects)
}

func (h *ProjectHandler) GetByID(w http.ResponseWriter, r *http.Request) {}

func (h *ProjectHandler) Update(w http.ResponseWriter, r *http.Request) {}

func (h *ProjectHandler) Delete() {}
