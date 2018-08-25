package service

import (
	"errors"
	"go_demo/common/util"
	"go_demo/model"
	"mime/multipart"
	"time"
)

func AddNewOrder(orderid string, username string, amount float64, status string, fileURL string) error {
	if orderid == "" || username == "" || amount == 0.0 || status == "" {
		return errors.New("username,amount,status can't be empty")
	}

	order := model.Order{OrderId: orderid, UserName: username, Amount: amount,
		Status: status, FileUrl: fileURL, CreateAt: time.Now()}

	return insertNewOrderRecord(&order)
}

func UpdateOrder(orderid string, amount float64, status string, fileURL string) error {
	if orderid == "" || amount == 0.0 || status == "" || fileURL == "" {
		return errors.New("username,amount,status can not be empty")
	}

	order := model.Order{OrderId: orderid, Amount: amount, Status: status, FileUrl: fileURL}
	return updateOrder(&order)
}

func GetOrder(rec *model.Order) error {
	return getOrder(rec)
}

func GetOrderListByUserName(key string, rec []model.Order) error {
	if key == "" {
		return errors.New("username can not be empty")
	}

	return getOrderListByUserName(key, rec)
}

func UpdateOrderFileURL(orderid string, fileURL string) error {
	if orderid == "" || fileURL == "" {
		return errors.New("order id ,file URL can not be empty")
	}

	order := model.Order{OrderId: orderid, FileUrl: fileURL}
	return updateOrderFileURL(&order)
}

func GetOrderFileURL(orderid string) (string, error) {
	if orderid == "" {
		return "", errors.New("order id can not be empty")
	}

	order := model.Order{OrderId: orderid}
	err := getOrder(&order)
	return order.FileUrl, err
}

func UploadFile(file multipart.File, orderid string, filedir string) (errR error) {
	// save file
	fileURL, err := util.SaveFile(file, filedir)
	if err != nil {
		return errors.New("save file error：" + err.Error())
	}

	// update DB file_url
	if err = UpdateOrderFileURL(orderid, fileURL); err != nil {
		return errors.New("update file URL error：" + err.Error())
	}
	return err
}
