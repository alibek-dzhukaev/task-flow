package handler

import (
	"encoding/json"
	"net/http"

	"github.com/alibek-dzhukaev/task-flow/internal/dto"
	"github.com/alibek-dzhukaev/task-flow/internal/service"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Fail(w, http.StatusBadRequest, "invalid request body")
		return
	}

	user, err := h.authService.Register(req.Email, req.Password, req.FirstName, req.LastName)
	if err != nil {
		Fail(w, http.StatusBadRequest, err.Error())
		return
	}

	Success(w, http.StatusCreated, user)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Fail(w, http.StatusBadRequest, "invalid request body")
		return
	}

	token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		Fail(w, http.StatusUnauthorized, err.Error())
		return
	}

	Success(w, http.StatusOK, dto.AuthResponse{AccessToken: token})
}
