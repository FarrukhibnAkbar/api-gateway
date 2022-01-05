package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/FarrukhibnAkbar/api-gateway/api/docs"
	v1 "github.com/FarrukhibnAkbar/api-gateway/api/handlers/v1"
	"github.com/FarrukhibnAkbar/api-gateway/config"
	"github.com/FarrukhibnAkbar/api-gateway/pkg/logger"
	"github.com/FarrukhibnAkbar/api-gateway/services"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")
	api.POST("/tasks", handlerV1.CreateTask)
	api.GET("/tasks/:id", handlerV1.GetTask)
	api.GET("/tasks", handlerV1.ListTasks)
	api.PUT("/tasks/:id", handlerV1.UpdateTask)
	api.DELETE("/tasks/:id", handlerV1.DeleteTask)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
