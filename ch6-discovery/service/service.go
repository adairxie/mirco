package service

import "context"

type Service interface {
	// 健康检查接口
	HealthCheck() bool
	// 打招呼接口
	SayHello() string
	// 服务发现接口
	DiscoveryService(ctx context.Context, serviceName string) ([]interface{}, error)
}

