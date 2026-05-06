package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/schulzdimitri/portfolio/backend/internal/domain"
	"github.com/schulzdimitri/portfolio/backend/internal/handler"
)

type mockRepo struct{ shouldFail bool }

func (m *mockRepo) Save(_ context.Context, _ domain.ContactMessage) error {
	if m.shouldFail {
		return errors.New("db error")
	}
	return nil
}

type mockSender struct{ shouldFail bool }

func (m *mockSender) Send(_ domain.ContactMessage) error {
	if m.shouldFail {
		return errors.New("smtp unavailable")
	}
	return nil
}

func postContact(t *testing.T, body any, repo handler.ContactRepository, sender handler.Sender) *httptest.ResponseRecorder {
	t.Helper()
	raw, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/api/contact", bytes.NewReader(raw))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ContactHandler(repo, sender)(rec, req)
	return rec
}

func TestContact_ValidPayload_ReturnsAccepted(t *testing.T) {
	body := domain.ContactMessage{Name: "Dimitri", Email: "d@example.com", Message: "Hello!"}
	rec := postContact(t, body, &mockRepo{}, &mockSender{})
	if rec.Code != http.StatusAccepted {
		t.Errorf("expected 202, got %d", rec.Code)
	}
}

func TestContact_MissingName_ReturnsUnprocessable(t *testing.T) {
	body := domain.ContactMessage{Email: "d@example.com", Message: "Hello!"}
	rec := postContact(t, body, &mockRepo{}, &mockSender{})
	if rec.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected 422, got %d", rec.Code)
	}
}

func TestContact_InvalidEmail_ReturnsUnprocessable(t *testing.T) {
	body := domain.ContactMessage{Name: "Dimitri", Email: "notanemail", Message: "Hello!"}
	rec := postContact(t, body, &mockRepo{}, &mockSender{})
	if rec.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected 422, got %d", rec.Code)
	}
}

func TestContact_NameWithOnlySpaces_ReturnsUnprocessable(t *testing.T) {
	body := domain.ContactMessage{Name: "   ", Email: "d@example.com", Message: "Hello!"}
	rec := postContact(t, body, &mockRepo{}, &mockSender{})
	if rec.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected 422 for whitespace-only name, got %d", rec.Code)
	}
}

func TestContact_RepoFailure_ReturnsInternalServerError(t *testing.T) {
	body := domain.ContactMessage{Name: "Dimitri", Email: "d@example.com", Message: "Hello!"}
	rec := postContact(t, body, &mockRepo{shouldFail: true}, &mockSender{})
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", rec.Code)
	}
}

func TestContact_SenderFailure_StillReturnsAccepted(t *testing.T) {
	// Email is best-effort — DB saved successfully, so 202 is expected even if email fails
	body := domain.ContactMessage{Name: "Dimitri", Email: "d@example.com", Message: "Hello!"}
	rec := postContact(t, body, &mockRepo{}, &mockSender{shouldFail: true})
	if rec.Code != http.StatusAccepted {
		t.Errorf("expected 202 even with sender failure, got %d", rec.Code)
	}
}

func TestContact_WrongMethod_ReturnsMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/contact", nil)
	rec := httptest.NewRecorder()
	handler.ContactHandler(&mockRepo{}, &mockSender{})(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", rec.Code)
	}
}
