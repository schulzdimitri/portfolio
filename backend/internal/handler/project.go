package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/schulzdimitri/portfolio/backend/internal/domain"
	"github.com/schulzdimitri/portfolio/backend/internal/repository"
)

type ProjectHandler struct {
	repo repository.ProjectRepository
}

func NewProjectHandler(repo repository.ProjectRepository) *ProjectHandler {
	return &ProjectHandler{
		repo: repo,
	}
}

func (h *ProjectHandler) GetProjects(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, `{"error": "method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	projects, err := h.repo.GetAll()
	if err != nil {
		slog.Error("Failed to fetch projects", "error", err)
		http.Error(w, `{"error": "failed to fetch projects"}`, http.StatusInternalServerError)
		return
	}

	// Wrapper object to match the expected format: {"projects": [...]}
	response := struct {
		Projects interface{} `json:"projects"`
	}{
		Projects: projects,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("Failed to encode projects response", "error", err)
	}
}

func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var p domain.Project
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, `{"error": "invalid json payload"}`, http.StatusBadRequest)
		return
	}

	if err := h.repo.Insert(&p); err != nil {
		slog.Error("Failed to insert project", "error", err)
		http.Error(w, `{"error": "failed to insert project"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		slog.Error("Failed to encode created project", "error", err)
	}
}
