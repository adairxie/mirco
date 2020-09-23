package service

import (
	"context"
	"discovery/config"
	"discovery/discover"
	"errors"
)

type Service interface {
	// 健康检查接口
	HealthCheck() bool
	// 打招呼接口
	SayHello() string
	// 服务发现接口
	DiscoveryService(ctx context.Context, serviceName string) ([]interface{}, error)
}

var ErrNotServiceInstances = errors.New("instances are not existed")

type DiscoveryServiceImpl struct {
	discoveryClient discover.DiscoveryClient
}

func NewDiscoveryServiceImpl(discoveryClient discover.DiscoveryClient) Service {
	return &DiscoveryServiceImpl{
		discoveryClient: discoveryClient,
	}
}

func (*DiscoveryServiceImpl) SayHello() string {
	return "Hello World!"
}

func (service *DiscoveryServiceImpl) DiscoveryService(ctx context.Context, serviceName string) ([]interface{}, error) {

	instances := service.discoveryClient.DiscoverServices(serviceName, config.Logger)
	return instances, nil
}

// HealthCheck implement Service method
func (*DiscoveryServiceImpl) HealthCheck() bool {
	return true
}
