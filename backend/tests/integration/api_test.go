package integration

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/schulzdimitri/portfolio/backend/internal/domain"
	"github.com/schulzdimitri/portfolio/backend/internal/handler"
	"github.com/schulzdimitri/portfolio/backend/internal/middleware"
	"github.com/schulzdimitri/portfolio/backend/internal/repository"
	"github.com/schulzdimitri/portfolio/backend/internal/sender"
)

func setupTestServer(t *testing.T) (*httptest.Server, *sql.DB) {
	tempDB := filepath.Join(t.TempDir(), "test_integration.db")
	db, err := repository.NewSQLiteDB(tempDB)
	if err != nil {
		t.Fatalf("failed to init db: %v", err)
	}

	contactRepo := repository.NewSQLiteContactRepository(db)
	projectRepo := repository.NewSQLiteProjectRepository(db)
	experienceRepo := repository.NewSQLiteExperienceRepository(db)

	emailSender := sender.NewSMTP(sender.SMTPConfig{})

	contactLimiter := middleware.NewRateLimiter(5, time.Minute)
	projectHandler := handler.NewProjectHandler(projectRepo)
	experienceHandler := handler.NewExperienceHandler(experienceRepo)

	adminToken := "test_token"
	authMiddleware := middleware.RequireAuth(adminToken)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", handler.Health)

	mux.HandleFunc("GET /api/projects", projectHandler.GetProjects)
	mux.Handle("POST /api/projects", authMiddleware(http.HandlerFunc(projectHandler.CreateProject)))

	mux.HandleFunc("GET /api/experiences", experienceHandler.GetExperiences)
	mux.Handle("POST /api/experiences", authMiddleware(http.HandlerFunc(experienceHandler.CreateExperience)))

	mux.Handle("/api/contact", contactLimiter.Middleware(handler.ContactHandler(contactRepo, emailSender)))

	ts := httptest.NewServer(mux)

	return ts, db
}

func TestHealthCheck(t *testing.T) {
	ts, db := setupTestServer(t)
	defer ts.Close()
	defer db.Close()

	resp, err := http.Get(ts.URL + "/api/health")
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestProjectIntegration(t *testing.T) {
	ts, db := setupTestServer(t)
	defer ts.Close()
	defer db.Close()

	resp, err := http.Post(ts.URL+"/api/projects", "application/json", bytes.NewBufferString(`{}`))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("expected unauthorized, got %d", resp.StatusCode)
	}

	proj := domain.Project{
		Title:       "Test Project",
		Description: "Integration test project",
		Github:      "https://github.com/test",
		Tags:        []string{"Go", "Test"},
	}
	payload, _ := json.Marshal(proj)

	req, _ := http.NewRequest("POST", ts.URL+"/api/projects", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer test_token")

	client := &http.Client{}
	resp2, err := client.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp2.Body.Close()
	if resp2.StatusCode != http.StatusCreated {
		t.Errorf("expected created 201, got %d", resp2.StatusCode)
	}

	resp3, err := http.Get(ts.URL + "/api/projects")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp3.Body.Close()
	if resp3.StatusCode != http.StatusOK {
		t.Errorf("expected ok 200, got %d", resp3.StatusCode)
	}

	var respObj struct {
		Projects []domain.Project `json:"projects"`
	}
	if err := json.NewDecoder(resp3.Body).Decode(&respObj); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	projects := respObj.Projects
	if len(projects) != 1 || projects[0].Title != "Test Project" {
		t.Errorf("unexpected project list: %+v", projects)
	}
}

func TestExperienceIntegration(t *testing.T) {
	ts, db := setupTestServer(t)
	defer ts.Close()
	defer db.Close()

	exp := domain.Experience{
		Company: "Test Company",
		Role:    "Tester",
		Period:  "2023-2024",
		Duties:  []string{"Testing stuff"},
	}
	payload, _ := json.Marshal(exp)

	req, _ := http.NewRequest("POST", ts.URL+"/api/experiences", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer test_token")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Errorf("expected created, got %d", res.StatusCode)
	}

	res2, err := http.Get(ts.URL + "/api/experiences")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer res2.Body.Close()

	var exps []domain.Experience
	if err := json.NewDecoder(res2.Body).Decode(&exps); err != nil {
		t.Fatalf("failed to decode: %v", err)
	}

	if len(exps) != 1 || exps[0].Company != "Test Company" {
		t.Errorf("unexpected experiences list: %+v", exps)
	}
}

func TestContactIntegration(t *testing.T) {
	ts, db := setupTestServer(t)
	defer ts.Close()
	defer db.Close()

	os.Setenv("TEST_NO_SMTP", "1")
	defer os.Unsetenv("TEST_NO_SMTP")

	contact := domain.ContactMessage{
		Name:    "John Doe",
		Email:   "john@example.com",
		Message: "Hello, this is a test.",
	}
	payload, _ := json.Marshal(contact)

	res, err := http.Post(ts.URL+"/api/contact", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusAccepted {
		t.Errorf("expected 202 Accepted, got %d", res.StatusCode)
	}
}
