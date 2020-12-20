package main

import (
	"fmt"

	"github.com/songdony/go_ioc/Injector"
	"github.com/songdony/go_ioc/services"
)

func main(){
	//uid := 1
	//userService := services.NewUserService(services.NewOrderService())
	//// 用户详细
	//userService.GetUserInfo(uid)
	//// 用户订单
	//userService.GetOrderInfo(uid)

	//Injector.BeanFactory.Set(services.NewOrderService())
	Injector.BeanFactory.Set(&services.OrderService{})
	order := Injector.BeanFactory.Get((*services.OrderService)(nil))
	fmt.Println(order)
}
