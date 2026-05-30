package handler

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Success(w http.ResponseWriter, status int, data any) {
	JSON(w, status, Response{
		Success: true,
		Data:    data,
	})
}

func Fail(w http.ResponseWriter, status int, message string) {
	JSON(w, status, ErrorResponse{
		Success: false,
		Error:   message,
	})
}
