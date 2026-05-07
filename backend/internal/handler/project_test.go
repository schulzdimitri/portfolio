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

type mockProjectRepo struct {
	projects []domain.Project
	err      error
}

func (m *mockProjectRepo) Insert(project *domain.Project) error {
	project.ID = len(m.projects) + 1
	m.projects = append(m.projects, *project)
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

	t.Run("POST Method Not Allowed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/projects", nil)
		w := httptest.NewRecorder()

		h.GetProjects(w, req)

		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("expected status 405, got %d", w.Code)
		}
	})

	t.Run("POST success", func(t *testing.T) {
		payload := `{"title":"Test Project", "description":"A project", "github":"link", "tags":["go"]}`
		req := httptest.NewRequest(http.MethodPost, "/api/projects", strings.NewReader(payload))
		w := httptest.NewRecorder()

		h.CreateProject(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("expected status 201, got %d", w.Code)
		}

		var resp domain.Project
		if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if resp.ID == 0 {
			t.Errorf("expected ID to be set, got 0")
		}
	})
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

func (m *mockProjectRepo) Delete(id int) error { return m.err }

func (m *mockProjectRepo) Update(id int, p *domain.Project) error { return m.err }

func TestUpdateProject(t *testing.T) {
	repo := &mockProjectRepo{}
	h := handler.NewProjectHandler(repo)

	t.Run("Method Not Allowed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/projects/1", nil)
		w := httptest.NewRecorder()
		h.UpdateProject(w, req)
		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("expected 405, got %d", w.Code)
		}
	})

	t.Run("Invalid ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/api/projects/abc", strings.NewReader(`{}`))
		req.SetPathValue("id", "abc")
		w := httptest.NewRecorder()
		h.UpdateProject(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", w.Code)
		}
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/api/projects/1", strings.NewReader(`{invalid`))
		req.SetPathValue("id", "1")
		w := httptest.NewRecorder()
		h.UpdateProject(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", w.Code)
		}
	})

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/api/projects/1", strings.NewReader(`{"title":"Updated"}`))
		req.SetPathValue("id", "1")
		w := httptest.NewRecorder()
		h.UpdateProject(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", w.Code)
		}
	})

	t.Run("Not Found", func(t *testing.T) {
		repo.err = errors.New("project not found")
		req := httptest.NewRequest(http.MethodPut, "/api/projects/1", strings.NewReader(`{"title":"Updated"}`))
		req.SetPathValue("id", "1")
		w := httptest.NewRecorder()
		h.UpdateProject(w, req)
		if w.Code != http.StatusNotFound {
			t.Errorf("expected 404, got %d", w.Code)
		}
	})

	t.Run("Internal Error", func(t *testing.T) {
		repo.err = errors.New("other error")
		req := httptest.NewRequest(http.MethodPut, "/api/projects/1", strings.NewReader(`{"title":"Updated"}`))
		req.SetPathValue("id", "1")
		w := httptest.NewRecorder()
		h.UpdateProject(w, req)
		if w.Code != http.StatusInternalServerError {
			t.Errorf("expected 500, got %d", w.Code)
		}
	})
}

func TestDeleteProject(t *testing.T) {
	repo := &mockProjectRepo{}
	h := handler.NewProjectHandler(repo)

	t.Run("Method Not Allowed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/projects/1", nil)
		w := httptest.NewRecorder()
		h.DeleteProject(w, req)
		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("expected 405, got %d", w.Code)
		}
	})

	t.Run("Invalid ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/projects/abc", nil)
		req.SetPathValue("id", "abc")
		w := httptest.NewRecorder()
		h.DeleteProject(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", w.Code)
		}
	})

	t.Run("Not Found", func(t *testing.T) {
		repo.err = errors.New("project not found")
		req := httptest.NewRequest(http.MethodDelete, "/api/projects/1", nil)
		req.SetPathValue("id", "1")
		w := httptest.NewRecorder()
		h.DeleteProject(w, req)
		if w.Code != http.StatusNotFound {
			t.Errorf("expected 404, got %d", w.Code)
		}
	})

	t.Run("Success", func(t *testing.T) {
		repo.err = nil
		req := httptest.NewRequest(http.MethodDelete, "/api/projects/1", nil)
		req.SetPathValue("id", "1")
		w := httptest.NewRecorder()
		h.DeleteProject(w, req)
		if w.Code != http.StatusNoContent {
			t.Errorf("expected 204, got %d", w.Code)
		}
	})
}
