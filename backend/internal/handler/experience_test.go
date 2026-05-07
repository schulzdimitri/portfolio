package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/schulzdimitri/portfolio/backend/internal/domain"
	"github.com/schulzdimitri/portfolio/backend/internal/handler"
)

type mockExperienceRepository struct {
	experiences []domain.Experience
	err         error
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

	t.Run("POST success", func(t *testing.T) {
		payload := `{"company":"Test Co", "role":"Dev", "period":"2024", "duties":["Coding"]}`
		req := httptest.NewRequest(http.MethodPost, "/api/experiences", strings.NewReader(payload))
		w := httptest.NewRecorder()

		h.CreateExperience(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("expected status 201, got %d", w.Code)
		}

		var resp domain.Experience
		if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if resp.ID == 0 {
			t.Errorf("expected ID to be set, got 0")
		}
	})
}

func (m *mockExperienceRepository) Delete(id int) error                       { return m.err }
func (m *mockExperienceRepository) Update(id int, p *domain.Experience) error { return m.err }

func TestUpdateExperience(t *testing.T) {
	repo := &mockExperienceRepository{}
	h := handler.NewExperienceHandler(repo)

	t.Run("Method Not Allowed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/experiences/1", nil)
		w := httptest.NewRecorder()
		h.UpdateExperience(w, req)
		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("expected 405, got %d", w.Code)
		}
	})

	t.Run("Invalid ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/api/experiences/abc", strings.NewReader(`{}`))
		req.SetPathValue("id", "abc")
		w := httptest.NewRecorder()
		h.UpdateExperience(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", w.Code)
		}
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/api/experiences/1", strings.NewReader(`{invalid`))
		req.SetPathValue("id", "1")
		w := httptest.NewRecorder()
		h.UpdateExperience(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", w.Code)
		}
	})

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/api/experiences/1", strings.NewReader(`{"company":"Updated"}`))
		req.SetPathValue("id", "1")
		w := httptest.NewRecorder()
		h.UpdateExperience(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", w.Code)
		}
	})

	t.Run("Not Found", func(t *testing.T) {
		repo.err = errors.New("experience not found")
		req := httptest.NewRequest(http.MethodPut, "/api/experiences/1", strings.NewReader(`{"company":"Updated"}`))
		req.SetPathValue("id", "1")
		w := httptest.NewRecorder()
		h.UpdateExperience(w, req)
		if w.Code != http.StatusNotFound {
			t.Errorf("expected 404, got %d", w.Code)
		}
	})

	t.Run("Internal Error", func(t *testing.T) {
		repo.err = errors.New("other error")
		req := httptest.NewRequest(http.MethodPut, "/api/experiences/1", strings.NewReader(`{"company":"Updated"}`))
		req.SetPathValue("id", "1")
		w := httptest.NewRecorder()
		h.UpdateExperience(w, req)
		if w.Code != http.StatusInternalServerError {
			t.Errorf("expected 500, got %d", w.Code)
		}
	})
}

func TestDeleteExperience(t *testing.T) {
	repo := &mockExperienceRepository{}
	h := handler.NewExperienceHandler(repo)

	t.Run("Method Not Allowed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/experiences/1", nil)
		w := httptest.NewRecorder()
		h.DeleteExperience(w, req)
		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("expected 405, got %d", w.Code)
		}
	})

	t.Run("Invalid ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/experiences/abc", nil)
		req.SetPathValue("id", "abc")
		w := httptest.NewRecorder()
		h.DeleteExperience(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", w.Code)
		}
	})

	t.Run("Not Found", func(t *testing.T) {
		repo.err = errors.New("experience not found")
		req := httptest.NewRequest(http.MethodDelete, "/api/experiences/1", nil)
		req.SetPathValue("id", "1")
		w := httptest.NewRecorder()
		h.DeleteExperience(w, req)
		if w.Code != http.StatusNotFound {
			t.Errorf("expected 404, got %d", w.Code)
		}
	})

	t.Run("Success", func(t *testing.T) {
		repo.err = nil
		req := httptest.NewRequest(http.MethodDelete, "/api/experiences/1", nil)
		req.SetPathValue("id", "1")
		w := httptest.NewRecorder()
		h.DeleteExperience(w, req)
		if w.Code != http.StatusNoContent {
			t.Errorf("expected 204, got %d", w.Code)
		}
	})
}
