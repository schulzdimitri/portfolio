package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/schulzdimitri/portfolio/backend/internal/handler"
)

type mockSender struct {
	shouldFail bool
}

func (m *mockSender) Send(_ handler.ContactRequest) error {
	if m.shouldFail {
		return errors.New("smtp unavailable")
	}
	return nil
}

func postContact(t *testing.T, body any, sender handler.Sender) *httptest.ResponseRecorder {
	t.Helper()
	raw, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/api/contact", bytes.NewReader(raw))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ContactHandler(sender)(rec, req)
	return rec
}

func TestContact_ValidPayload_ReturnsAccepted(t *testing.T) {
	body := handler.ContactRequest{Name: "Dimitri", Email: "d@example.com", Message: "Hello!"}
	rec := postContact(t, body, &mockSender{})

	if rec.Code != http.StatusAccepted {
		t.Errorf("expected 202, got %d", rec.Code)
	}
}

func TestContact_MissingName_ReturnsUnprocessable(t *testing.T) {
	body := handler.ContactRequest{Email: "d@example.com", Message: "Hello!"}
	rec := postContact(t, body, &mockSender{})

	if rec.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected 422, got %d", rec.Code)
	}
}

func TestContact_InvalidEmail_ReturnsUnprocessable(t *testing.T) {
	body := handler.ContactRequest{Name: "Dimitri", Email: "notanemail", Message: "Hello!"}
	rec := postContact(t, body, &mockSender{})

	if rec.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected 422, got %d", rec.Code)
	}
}

func TestContact_SenderFailure_ReturnsInternalServerError(t *testing.T) {
	body := handler.ContactRequest{Name: "Dimitri", Email: "d@example.com", Message: "Hello!"}
	rec := postContact(t, body, &mockSender{shouldFail: true})

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", rec.Code)
	}
}

func TestContact_WrongMethod_ReturnsMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/contact", nil)
	rec := httptest.NewRecorder()
	handler.ContactHandler(&mockSender{})(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", rec.Code)
	}
}
