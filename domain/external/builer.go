package external

import (
	"sofia-backend/config"
	externalservices "sofia-backend/infraestructure/external-services"
)

type ServiceContainer struct {
	EmailService *EmailService
}

func Build(
	external *externalservices.ExternalServicesContainer,
	config *config.Config,
) *ServiceContainer {
	return &ServiceContainer{
		EmailService: NewEmailService(external.MailerService, external.RenderService, config),
	}
}
