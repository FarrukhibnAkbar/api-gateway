package v1

import (
	"github.com/FarrukhibnAkbar/api-gateway/config"
	"github.com/FarrukhibnAkbar/api-gateway/pkg/logger"
	"github.com/FarrukhibnAkbar/api-gateway/services"
)

type handlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
}

type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
}

func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
	}
}
