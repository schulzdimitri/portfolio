package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/schulzdimitri/portfolio/backend/internal/domain"
	"github.com/schulzdimitri/portfolio/backend/internal/handler"
)

type mockProjectRepo struct {
	projects []domain.Project
	err      error
}

func (m *mockProjectRepo) Insert(project *domain.Project) error {
	return nil
}

func (m *mockProjectRepo) GetAll() ([]domain.Project, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.projects, nil
}

func (m *mockProjectRepo) Count() (int, error) {
	return len(m.projects), nil
}

func TestGetProjects_Success(t *testing.T) {
	repo := &mockProjectRepo{
		projects: []domain.Project{
			{ID: 1, Title: "A"},
			{ID: 2, Title: "B"},
		},
	}
	h := handler.NewProjectHandler(repo)

	req := httptest.NewRequest(http.MethodGet, "/api/projects", nil)
	w := httptest.NewRecorder()

	h.GetProjects(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var res struct {
		Projects []domain.Project `json:"projects"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}

	if len(res.Projects) != 2 {
		t.Errorf("expected 2 projects, got %d", len(res.Projects))
	}
}

func TestGetProjects_MethodNotAllowed(t *testing.T) {
	repo := &mockProjectRepo{}
	h := handler.NewProjectHandler(repo)

	req := httptest.NewRequest(http.MethodPost, "/api/projects", nil)
	w := httptest.NewRecorder()

	h.GetProjects(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status 405, got %d", w.Code)
	}
}

func TestGetProjects_RepoError(t *testing.T) {
	repo := &mockProjectRepo{
		err: errors.New("db down"),
	}
	h := handler.NewProjectHandler(repo)

	req := httptest.NewRequest(http.MethodGet, "/api/projects", nil)
	w := httptest.NewRecorder()

	h.GetProjects(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", w.Code)
	}
}
