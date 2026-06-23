package facades

import (
	"context"
	"fmt"
	"sofia-backend/domain/external"
	"sofia-backend/domain/services"
	"time"
)

func sendSupplierViewTemplateEmail(ctx context.Context, brandingService *services.CompanyBrandingService, emailService *external.EmailService, companyID string, to string, token string, exp time.Time) {
	if to == "" {
		return
	}
	if brandingService == nil {
		if err := emailService.SendSupplierViewEmail(to, token, exp); err != nil {
			fmt.Println("Error sending fallback supplier email:", err)
		}
		return
	}

	url := emailService.SupplierViewURL(token)
	subject, body, err := brandingService.ResolveEmailTemplate(ctx, companyID, "supplier_view", map[string]string{
		"url":            url,
		"expirationDate": exp.Format("01-02-2006 15:04 MST"),
	})
	if err == nil {
		err = emailService.SendHTML(to, subject, body)
	}
	if err != nil {
		fmt.Println("Error sending templated supplier email:", err)
		if fallbackErr := emailService.SendSupplierViewEmail(to, token, exp); fallbackErr != nil {
			fmt.Println("Error sending fallback supplier email:", fallbackErr)
		}
	}
}
