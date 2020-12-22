package main

import (
	"fmt"

	"github.com/songdony/go_ioc/Config"
	. "github.com/songdony/go_ioc/Injector"
	"github.com/songdony/go_ioc/services"
)

func main(){
	serviceConfig := Config.NewServiceConfig()
	BeanFactory.Config(serviceConfig)

	{
		// 这里测试 userService
		userService := services.NewUserService()
		BeanFactory.Apply(userService)
		fmt.Println(userService.Order)
	}
	{
		// 这里测试 adminService
		adminService := services.NewAdminService()
		BeanFactory.Apply(adminService)
		fmt.Println(adminService.Order)
	}
}
