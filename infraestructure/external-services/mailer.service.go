package externalservices

import (
	"fmt"
	"sofia-backend/config"
	"sofia-backend/domain/ports"
	"sofia-backend/types"

	"github.com/resend/resend-go/v2"
)

type MailerService struct {
	resend  *resend.Client
	from    string
	isDebug bool
}

func NewMailerService(cfg *config.Config) ports.PortMailer {
	resendClient := resend.NewClient(cfg.Mailer.ApiKey)
	return &MailerService{resend: resendClient, from: cfg.Mailer.MailFrom, isDebug: cfg.Debug}
}

func (m *MailerService) Send(to []string, subject string, body string) error {

	if m.isDebug {
		fmt.Printf("Debug: Sending email to %v with subject '%s'\n", to, subject)
		fmt.Printf("Debug: Email body: %s\n", body)
		return nil // Skip actual sending in debug mode
	}

	params := &resend.SendEmailRequest{
		From:    m.from,
		To:      to,
		Subject: subject,
		Html:    body,
	}

	_, err := m.resend.Emails.Send(params)
	if err != nil {
		return types.ThrowData(fmt.Sprintf("Error al enviar el email: %v", err))
	}

	return nil
}
