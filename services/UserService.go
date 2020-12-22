package services

import "fmt"

type UserService struct {
	Order *OrderService `inject:"ServiceConfig.OrderService()"`   // -
}

func NewUserService() *UserService {
	return &UserService{}
}

func(this *UserService) GetUserInfo(uid int){
	fmt.Println("获取用户ID=",uid,"的详细信息")
}

func(this *UserService) GetOrderInfo(uid int){
	NewOrderService().GetOrderInfo(uid)
}