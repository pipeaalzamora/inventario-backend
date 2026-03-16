package externalservices

import (
	"sofia-backend/config"
	"sofia-backend/domain/ports"

	"github.com/redis/go-redis/v9"
)

type ExternalServicesContainer struct {
	RenderService ports.PortRender
	MailerService ports.PortMailer
	CacheService  ports.PortCache
	SearchService ports.PortSearch
	SSEservice    ports.PortSee
	BucketService ports.PortBucket
}

func Build(
	cfg *config.Config,
	redisClient *redis.Client,
) *ExternalServicesContainer {
	return &ExternalServicesContainer{
		RenderService: NewRendererService(),
		MailerService: NewMailerService(cfg),
		CacheService:  NewCacheService(redisClient),
		SearchService: NewSearchService(cfg),
		SSEservice:    NewSSEService(),
		BucketService: NewBucketService(cfg),
	}
}
