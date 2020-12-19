package main

import "github.com/songdony/go_ioc/services"

func main(){
	uid := 1
	userService := services.NewUserService(services.NewOrderService())
	// 用户详细
	userService.GetUserInfo(uid)
	// 用户订单
	userService.GetOrderInfo(uid)
}
