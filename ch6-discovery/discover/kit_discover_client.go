package discover

import (
	"log"
	"strconv"
	"sync"

	"github.com/hashicorp/consul/api/watch"

	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
)

// KitDiscoverClient is kit service discovery client
type KitDiscoverClient struct {
	Host         string // Consul Host
	Port         int    // Consul Port
	client       consul.Client
	config       *api.Config
	mutex        sync.Mutex
	instancesMap sync.Map
}

// NewKitDiscoverClient create an KitDiscoverClient instance
func NewKitDiscoverClient(consulHost string, consulPort int) (DiscoveryClient, error) {
	// 通过Consul Host 和 Consul Port 创建一个consul.Client
	consulConfig := api.DefaultConfig()
	consulConfig.Address = consulHost + ":" + strconv.Itoa(consulPort)
	apiClient, err := api.NewClient(consulConfig)
	if err != nil {
		return nil, err
	}
	client := consul.NewClient(apiClient)
	return &KitDiscoverClient{
		Host:   consulHost,
		Port:   consulPort,
		config: consulConfig,
		client: client,
	}, err
}

// Register register service meta to consul by consulClient
func (consulClient *KitDiscoverClient) Register(serviceName, instanceId,
	healthCheckUrl string, instanceHost string, instancePort int, meta map[string]string, logger *log.Logger) bool {
	// 1. 构建服务实例元数据
	serviceRegistration := &api.AgentServiceRegistration{
		ID:      instanceId,
		Name:    serviceName,
		Address: instanceHost,
		Port:    instancePort,
		Meta:    meta,
		Check: &api.AgentServiceCheck{
			DeregisterCriticalServiceAfter: "30s",
			HTTP:                           "http://" + instanceHost + ":" + strconv.Itoa(instancePort) + healthCheckUrl,
			Interval:                       "15s",
		},
	}

	// 2. 发送服务注册到Consul中
	err := consulClient.client.Register(serviceRegistration)
	if err != nil {
		log.Println("Register Service Error!")
		return false
	}
	log.Println("Register Service Success!")
	return true
}

// DeRegister deregister service by instanceId
func (consulClient *KitDiscoverClient) DeRegister(instanceId string, logger *log.Logger) bool {
	// 构建包含服务实例ID的元数据结构体
	serviceRegistration := &api.AgentServiceRegistration{
		ID: instanceId,
	}
	// 发送服务注销请求
	err := consulClient.client.Deregister(serviceRegistration)
	if err != nil {
		logger.Println("Deregister Service Error!")
		return false
	}
	log.Println("Deregister Service Success!")

	return true
}

// DiscoverServices find  services
func (consulClient *KitDiscoverClient) DiscoverServices(serviceName string, logger *log.Logger) []interface{} {

	instancList, ok := consulClient.instancesMap.Load(serviceName)
	if ok {
		logger.Println("DiscoverServices get cache service info")
		return instancList.([]interface{})
	}

	consulClient.mutex.Lock()
	instancList, ok = consulClient.instancesMap.Load(serviceName)
	if ok {
		return instancList.([]interface{})
	} else {
		go func() {
			params := make(map[string]interface{})
			params["type"] = "service"
			params["service"] = serviceName
			plan, _ := watch.Parse(params)
			plan.Handler = func(u uint64, i interface{}) {
				if i == nil {
					return
				}
				v, ok := i.([]*api.ServiceEntry)
				if !ok {
					return
				}
				if len(v) == 0 {
					consulClient.instancesMap.Store(serviceName, []interface{}{})
				} else {
					var healthServices []interface{}
					for _, service := range v {
						if service.Checks.AggregatedStatus() == api.HealthPassing {
							healthServices = append(healthServices, service.Service)
						}
					}
					consulClient.instancesMap.Store(serviceName, healthServices)
				}
			}
			defer plan.Stop()
			plan.Run(consulClient.config.Address)
		}()
	}
	defer consulClient.mutex.Unlock()

	entries, _, err := consulClient.client.Service(serviceName, "", false, nil)
	if err != nil {
		consulClient.instancesMap.Store(serviceName, []interface{}{})
		log.Println("Discover Service Error!")
		return nil
	}

	instances := make([]interface{}, len(entries))
	for i := 0; i < len(instances); i++ {
		instances[i] = entries[i].Service
	}

	consulClient.instancesMap.Store(serviceName, instances)
	return instances
}
