package services

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	"github.com/FarrukhibnAkbar/api-gateway/config"
	pb "github.com/FarrukhibnAkbar/api-gateway/genproto"
)

type IServiceManager interface {
	TaskService() pb.TaskServiceClient
}

type serviceManager struct {
	taskService pb.TaskServiceClient
}

func (s *serviceManager) TaskService() pb.TaskServiceClient {
	return s.taskService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connTask, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.TaskServiceHost, conf.TaskServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		taskService: pb.NewTaskServiceClient(connTask),
	}

	return serviceManager, nil
}
