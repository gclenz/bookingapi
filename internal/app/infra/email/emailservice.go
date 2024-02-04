package email

import (
	"bytes"
	"context"
	"net/smtp"
)

type IEmailService interface {
	Send(email string, content bytes.Buffer, ctx context.Context) error
}

func NewEmailService(username, password, host string) *EmailService {
	return &EmailService{
		Username: username,
		Password: password,
		Host:     host,
	}
}

type EmailService struct {
	Username string
	Password string
	Host     string
}

func (em *EmailService) Send(email string, content bytes.Buffer, ctx context.Context) error {
	auth := smtp.PlainAuth("", em.Username, em.Password, em.Host)

	to := []string{email}

	err := smtp.SendMail(em.Host+":587", auth, "no-reply@application.com", to, content.Bytes())
	return err
}
