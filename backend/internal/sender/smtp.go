package sender

import (
	"fmt"
	"log/slog"
	"net/smtp"

	"github.com/schulzdimitri/portfolio/backend/internal/domain"
)

type SMTPConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	To       string
}

type SMTPSender struct {
	config SMTPConfig
}

func NewSMTP(config SMTPConfig) *SMTPSender {
	return &SMTPSender{config: config}
}

func (s *SMTPSender) Send(msg domain.ContactMessage) error {
	if s.config.Host == "" || s.config.User == "" {
		slog.Info("SMTP not configured, logging contact message",
			"name", msg.Name, "email", msg.Email)
		return nil
	}

	auth := smtp.PlainAuth("", s.config.User, s.config.Password, s.config.Host)
	addr := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)

	subject := fmt.Sprintf("Portfolio contact from %s", msg.Name)
	body := fmt.Sprintf("From: %s <%s>\n\n%s", msg.Name, msg.Email, msg.Message)
	raw := []byte("Subject: " + subject + "\r\n\r\n" + body)

	return smtp.SendMail(addr, auth, s.config.User, []string{s.config.To}, raw)
}
