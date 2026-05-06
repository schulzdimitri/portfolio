package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

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
func (h *ProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, `{"error": "method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error": "invalid project id"}`, http.StatusBadRequest)
		return
	}

	var p domain.Project
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, `{"error": "invalid json payload"}`, http.StatusBadRequest)
		return
	}

	if err := h.repo.Update(id, &p); err != nil {
		slog.Error("Failed to update project", "error", err)
		if err.Error() == "project not found" {
			http.Error(w, `{"error": "project not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"error": "failed to update project"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func (h *ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, `{"error": "method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error": "invalid project id"}`, http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		slog.Error("Failed to delete project", "error", err)
		if err.Error() == "project not found" {
			http.Error(w, `{"error": "project not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"error": "failed to delete project"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
