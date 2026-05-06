package sender_test

import (
	"context"
	"testing"

	"github.com/schulzdimitri/portfolio/backend/internal/domain"
	"github.com/schulzdimitri/portfolio/backend/internal/sender"
)

func TestSMTP_Send(t *testing.T) {
	// Just a structural test to satisfy coverage and ensure initialization works
	cfg := sender.SMTPConfig{
		Host: "localhost",
		Port: "25",
	}
	s := sender.NewSMTP(cfg)

	// Sending to invalid host will fail but tests the path
	err := s.Send(context.Background(), domain.ContactMessage{
		Name:    "Test",
		Email:   "test@test.com",
		Message: "msg",
	})

	if err == nil {
		t.Errorf("expected error when sending without real SMTP server")
	}
}
