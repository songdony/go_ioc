package Config

import (
	"log"

	"github.com/songdony/go_ioc/services"
)

type ServiceConfig struct{

}

func NewServiceConfig() *ServiceConfig{
	return &ServiceConfig{}
}

func(this *ServiceConfig) OrderService() *services.OrderService{
	log.Println("初始化orderservice")
	return services.NewOrderService()
}
