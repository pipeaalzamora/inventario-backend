package external

import (
	"fmt"
	"html"
	"sofia-backend/config"
	"sofia-backend/domain/ports"
	"time"
)

type EmailService struct {
	emailService  ports.PortMailer
	renderService ports.PortRender
	frontUrl      string
}

func NewEmailService(emailService ports.PortMailer, renderService ports.PortRender, config *config.Config) *EmailService {
	return &EmailService{
		emailService:  emailService,
		renderService: renderService,
		frontUrl:      config.FrontUrl,
	}
}

func (m *EmailService) SendRecoveryEmail(to string, code string, ttl time.Time) error {
	// send the recovery code to the user via email
	html, err := m.renderService.Render("templates/recovery_code.html", map[string]interface{}{
		"Code": code,
		"TTL":  ttl,
	})
	if err != nil {
		return err
	}

	err = m.emailService.Send([]string{to}, "Recuperación de contraseña", html)
	if err != nil {
		return err
	}

	return nil
}

func (m *EmailService) SendWelcomeEmail(to string, generatedPassword string, url string) error {
	// send the welcome email to the user
	html, err := m.renderService.Render("templates/welcome.html", map[string]interface{}{
		"Password": generatedPassword,
		"Url":      url,
	})
	if err != nil {
		return err
	}

	err = m.emailService.Send([]string{to}, "Bienvenido a Sofia", html)
	if err != nil {
		return err
	}

	return nil
}

func (m *EmailService) SendSupplierViewEmail(to string, token string, exp time.Time) error {
	url := fmt.Sprintf("%s/supplier-manage-oc?token=%s", m.frontUrl, token)
	html, err := m.renderService.Render("templates/supplier_view.html", map[string]interface{}{
		"Url": html.EscapeString(url),
		"Exp": exp.Format("01-02-2006 15:04 MST"),
	})
	if err != nil {
		return err
	}

	err = m.emailService.Send([]string{to}, "Aprobar solicitud de compra", html)
	if err != nil {
		return err
	}

	return nil
}
