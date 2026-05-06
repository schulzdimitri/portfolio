package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type ContactRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type contactResponse struct {
	Message string `json:"message"`
}

type Sender interface {
	Send(req ContactRequest) error
}

func ContactHandler(sender Sender) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}

		var body ContactRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json body"})
			return
		}

		if err := validateContact(body); err != nil {
			writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
			return
		}

		if err := sender.Send(body); err != nil {
			log.Printf("sender error: %v", err)
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to process contact"})
			return
		}

		writeJSON(w, http.StatusAccepted, contactResponse{Message: "message received, thank you!"})
	}
}

func validateContact(r ContactRequest) error {
	r.Name = strings.TrimSpace(r.Name)
	r.Email = strings.TrimSpace(r.Email)
	r.Message = strings.TrimSpace(r.Message)

	switch {
	case r.Name == "":
		return fmt.Errorf("name is required")
	case r.Email == "" || !strings.Contains(r.Email, "@"):
		return fmt.Errorf("a valid email is required")
	case r.Message == "":
		return fmt.Errorf("message is required")
	}
	return nil
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}
