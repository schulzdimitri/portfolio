package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/schulzdimitri/portfolio/backend/internal/domain"
	"github.com/schulzdimitri/portfolio/backend/internal/handler"
)

type mockExperienceRepository struct {
	experiences []domain.Experience
}

func (m *mockExperienceRepository) Insert(exp *domain.Experience) error {
	exp.ID = len(m.experiences) + 1
	m.experiences = append(m.experiences, *exp)
	return nil
}

func (m *mockExperienceRepository) GetAll() ([]domain.Experience, error) {
	if m.experiences == nil {
		return make([]domain.Experience, 0), nil
	}
	return m.experiences, nil
}

func (m *mockExperienceRepository) Count() (int, error) {
	return len(m.experiences), nil
}

func TestExperienceHandler_GetExperiences(t *testing.T) {
	repo := &mockExperienceRepository{
		experiences: []domain.Experience{
			{ID: 1, Company: "Company A", Role: "Dev", Period: "2024", Duties: []string{"Coding"}},
		},
	}
	h := handler.NewExperienceHandler(repo)

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/experiences", nil)
		w := httptest.NewRecorder()

		h.GetExperiences(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}

		var resp []domain.Experience
		if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if len(resp) != 1 {
			t.Errorf("expected 1 experience, got %d", len(resp))
		}
	})

	t.Run("wrong method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/experiences", nil)
		w := httptest.NewRecorder()

		h.GetExperiences(w, req)

		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("expected status 405, got %d", w.Code)
		}
	})
}
