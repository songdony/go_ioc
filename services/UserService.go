package services

import "fmt"

type UserService struct {
	order *OrderService
}

func NewUserService(order *OrderService) *UserService {
	return &UserService{order: order}
}

func(this *UserService) GetUserInfo(uid int){
	fmt.Println("获取用户ID=",uid,"的详细信息")
}

func(this *UserService) GetOrderInfo(uid int){
	NewOrderService().GetOrderInfo(uid)
}