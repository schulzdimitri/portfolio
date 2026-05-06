package sender

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/schulzdimitri/portfolio/backend/internal/handler"
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

func (s *SMTPSender) Send(req handler.ContactRequest) error {
	if s.config.Host == "" || s.config.User == "" {
		log.Printf("[contact] SMTP not configured — logging message from %s <%s>: %s",
			req.Name, req.Email, req.Message)
		return nil
	}

	auth := smtp.PlainAuth("", s.config.User, s.config.Password, s.config.Host)
	addr := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)

	subject := fmt.Sprintf("Portfolio contact from %s", req.Name)
	body := fmt.Sprintf("From: %s <%s>\n\n%s", req.Name, req.Email, req.Message)
	msg := []byte("Subject: " + subject + "\r\n\r\n" + body)

	return smtp.SendMail(addr, auth, s.config.User, []string{s.config.To}, msg)
}
