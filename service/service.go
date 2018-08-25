package service

import (
	"time"
	"go_demo/model"
)

func AddNewOrder(username string, amount float64, status string, fileURL string) error {
	orderid := NewOrderID()
	order := model.TOrder{OrderId: orderid, UserName: username, Amount: amount,
	Status: status, FileUrl: fileURL, CreateAt: time.Now()}

	return insertNewOrderRecord(&order)
}

func UpdateOrder(orderid string, amount float64, status string, fileURL string) error {
	order := model.TOrder{OrderId: orderid, Amount: amount, Status: status, FileUrl: fileURL}
	return updateOrder(&order)
}

func GetOrder(rec *model.TOrder) error {
	return getOrder(rec)
}

func GetOrderListByUserName(key string, rec []model.TOrder) error {
	return getOrderListByUserName(key, rec)
}

func UpdateOrderFileURL(orderid string, fileURL string) error {
	order := model.TOrder{OrderId:orderid, FileUrl:fileURL}
	return updateOrderFileURL(&order)
}

func GetOrderFilePath(orderid string) error {

}