package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/schulzdimitri/portfolio/backend/internal/domain"
)

type ContactRepository interface {
	Save(ctx context.Context, msg domain.ContactMessage) error
}

type Sender interface {
	Send(msg domain.ContactMessage) error
}

func ContactHandler(repo ContactRepository, sender Sender) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeJSON(w, http.StatusMethodNotAllowed, errResp("method not allowed"))
			return
		}

		var msg domain.ContactMessage
		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
			writeJSON(w, http.StatusBadRequest, errResp("invalid json body"))
			return
		}

		if err := validateContact(&msg); err != nil {
			writeJSON(w, http.StatusUnprocessableEntity, errResp(err.Error()))
			return
		}

		if err := repo.Save(r.Context(), msg); err != nil {
			slog.Error("failed to save contact to db", "error", err)
			writeJSON(w, http.StatusInternalServerError, errResp("failed to process contact"))
			return
		}

		if err := sender.Send(msg); err != nil {
			slog.Warn("email notification failed, message saved in db", "error", err)
		}

		writeJSON(w, http.StatusAccepted, map[string]string{"message": "message received, thank you!"})
	}
}

func validateContact(msg *domain.ContactMessage) error {
	msg.Name = strings.TrimSpace(msg.Name)
	msg.Email = strings.TrimSpace(msg.Email)
	msg.Message = strings.TrimSpace(msg.Message)

	switch {
	case msg.Name == "":
		return fmt.Errorf("name is required")
	case msg.Email == "" || !strings.Contains(msg.Email, "@"):
		return fmt.Errorf("a valid email is required")
	case msg.Message == "":
		return fmt.Errorf("message is required")
	}
	return nil
}

func errResp(msg string) map[string]string {
	return map[string]string{"error": msg}
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}
