package service

import (
	"errors"
	"go_demo/common/util"
	"go_demo/model"
	"mime/multipart"
	"time"
)

func AddNewOrder(username string, amount float64, status string, fileURL string) error {
	orderid, err := util.NewOrderID()
	if err != nil {
		return errors.New("add new order error" + err.Error())
	}
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
	order := model.TOrder{OrderId: orderid, FileUrl: fileURL}
	return updateOrderFileURL(&order)
}

func GetOrderFileURL(orderid string) (string, error) {
	order := model.TOrder{OrderId: orderid}
	err := getOrder(&order)
	return order.FileUrl, err
}

func SaveFile(file multipart.File, filepath string) (fileURL string, errR error) {
	return util.SaveFile(file, filepath)
}
