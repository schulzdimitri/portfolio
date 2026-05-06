package handler

import (
	"encoding/json"
	"net/http"
        "strconv"
        "log/slog"
	"github.com/schulzdimitri/portfolio/backend/internal/domain"
	"github.com/schulzdimitri/portfolio/backend/internal/repository"

)

type ExperienceHandler struct {
	repo repository.ExperienceRepository
}

func NewExperienceHandler(repo repository.ExperienceRepository) *ExperienceHandler {
	return &ExperienceHandler{repo: repo}
}

func (h *ExperienceHandler) GetExperiences(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	experiences, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, "Failed to get experiences", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(experiences); err != nil {
		http.Error(w, "Failed to encode experiences", http.StatusInternalServerError)
	}
}

func (h *ExperienceHandler) CreateExperience(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var exp domain.Experience
	if err := json.NewDecoder(r.Body).Decode(&exp); err != nil {
		http.Error(w, `{"error": "invalid json payload"}`, http.StatusBadRequest)
		return
	}

	if err := h.repo.Insert(&exp); err != nil {
		http.Error(w, `{"error": "failed to insert experience"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(exp); err != nil {
		http.Error(w, "Failed to encode created experience", http.StatusInternalServerError)
	}
}
func (h *ExperienceHandler) UpdateExperience(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPut {
                http.Error(w, `{"error": "method not allowed"}`, http.StatusMethodNotAllowed)
                return
        }

        idStr := r.PathValue("id")
        id, err := strconv.Atoi(idStr)
        if err != nil {
                http.Error(w, `{"error": "invalid experience id"}`, http.StatusBadRequest)
                return
        }

        var exp domain.Experience
        if err := json.NewDecoder(r.Body).Decode(&exp); err != nil {
                http.Error(w, `{"error": "invalid json payload"}`, http.StatusBadRequest)
                return
        }

        if err := h.repo.Update(id, &exp); err != nil {
                slog.Error("Failed to update experience", "error", err)
                if err.Error() == "experience not found" {
                        http.Error(w, `{"error": "experience not found"}`, http.StatusNotFound)
                        return
                }
                http.Error(w, `{"error": "failed to update experience"}`, http.StatusInternalServerError)
                return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(exp)
}

func (h *ExperienceHandler) DeleteExperience(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodDelete {
                http.Error(w, `{"error": "method not allowed"}`, http.StatusMethodNotAllowed)
                return
        }

        idStr := r.PathValue("id")
        id, err := strconv.Atoi(idStr)
        if err != nil {
                http.Error(w, `{"error": "invalid experience id"}`, http.StatusBadRequest)
                return
        }

        if err := h.repo.Delete(id); err != nil {
                slog.Error("Failed to delete experience", "error", err)
                if err.Error() == "experience not found" {
                        http.Error(w, `{"error": "experience not found"}`, http.StatusNotFound)
                        return
                }
                http.Error(w, `{"error": "failed to delete experience"}`, http.StatusInternalServerError)
                return
        }

        w.WriteHeader(http.StatusNoContent)
}