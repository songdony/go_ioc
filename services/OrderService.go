package services

import "fmt"

type OrderService struct {

}

func NewOrderService() *OrderService{
	return &OrderService{}
}

func(this *OrderService) GetOrderInfo(uid int){
	fmt.Println("获取用户ID=",uid,"的订单信息")
}
