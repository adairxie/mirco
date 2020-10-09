package main

import (
	"context"
	"discovery/config"
	"discovery/discover"
	"discovery/endpoint"
	"discovery/service"
	"discovery/transport"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	uuid "github.com/satori/go.uuid"
)

func main() {

	var (
		servicePort = flag.Int("service.port", 10086, "service port")
		serviceHost = flag.String("service.host", "127.0.0.1", "service host")
		serviceName = flag.String("service.name", "SayHello", "service name")
		// consul address
		consulPort = flag.Int("consul.port", 8500, "consul port")
		consulHost = flag.String("consul.host", "127.0.0.1", "consul host")
	)

	flag.Parse()

	ctx := context.Background()
	errChan := make(chan error)

	var discoveryClient discover.DiscoveryClient

	discoveryClient, err := discover.NewKitDiscoverClient(*consulHost, *consulPort)
	if err != nil {
		config.Logger.Println("Get Consul client failed")
		os.Exit(-1)
	}

	var svc = service.NewDiscoveryServiceImpl(discoveryClient)

	sayHelloEndpoint := endpoint.MakeSayHelloEndpoint(svc)
	discoveryEndpoint := endpoint.MakeDiscoveryEndpoint(svc)
	healthEndpoint := endpoint.MakeHealthCheckEnpoint(svc)

	endpts := endpoint.DiscoveryEndpoints{
		SayHelloEndpoint:    sayHelloEndpoint,
		DiscoveryEndpoint:   discoveryEndpoint,
		HealthCheckEndpoint: healthEndpoint,
	}

	//创建http.Handler
	r := transport.MakeHttpHandler(ctx, endpts, config.KitLogger)

	//定义服务实例ID
	instanceId := *serviceName + "-" + uuid.NewV4().String()

	//启动 http server
	go func() {
		config.Logger.Println("HTTP Server start at port:" + strconv.Itoa(*servicePort))
		//启动前执行注册
		if !discoveryClient.Register(*serviceName, instanceId, "/health", *serviceHost, *servicePort, nil, config.Logger) {
			config.Logger.Printf("string-service for service %s failed.", serviceName)
			os.Exit(-1)
		}
		handler := r
		errChan <- http.ListenAndServe(":"+strconv.Itoa(*servicePort), handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	error := <-errChan
	discoveryClient.DeRegister(instanceId, config.Logger)
	config.Logger.Println(error)
}
