package sender_test

import (
	"testing"

	"github.com/schulzdimitri/portfolio/backend/internal/domain"
	"github.com/schulzdimitri/portfolio/backend/internal/sender"
)

func TestSMTP_Send(t *testing.T) {
	cfg := sender.SMTPConfig{
		Host: "localhost",
		Port: "25",
	}
	s := sender.NewSMTP(cfg)

	err := s.Send(domain.ContactMessage{
		Name:    "Test",
		Email:   "test@test.com",
		Message: "msg",
	})

	if err == nil {
		t.Errorf("expected error when sending without real SMTP server")
	}
}
