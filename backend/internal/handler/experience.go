package handler

import (
	"encoding/json"
	"net/http"

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
