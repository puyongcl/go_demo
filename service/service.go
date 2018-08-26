package service

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
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

func ExportOrderListWithExcel() (excelFileURL string, errR error) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		return "", err
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "order_id"
	cell = row.AddCell()
	cell.Value = "user_name"
	cell = row.AddCell()
	cell.Value = "amount"
	cell = row.AddCell()
	cell.Value = "status"
	cell = row.AddCell()
	cell.Value = "file_url"

	//
	var rec []model.Order
	err = getOrderList(rec)
	if err != nil {
		return "", err
	}

	for _, order := range rec {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = order.OrderId //"order_id"
		cell = row.AddCell()
		cell.Value = order.UserName //"user_name"
		cell = row.AddCell()
		cell.Value = fmt.Sprintf("%v", order.Amount) //"amount"
		cell = row.AddCell()
		cell.Value = order.Status //"status"
		cell = row.AddCell()
		cell.Value = order.FileUrl //"file_url"
	}

	filepath, err := util.NewUUID()
	if err != nil {
		return "", errors.New("gen file name error:" + err.Error())
	}
	filepath += ".xlsx"
	filepath = FileDir + filepath
	err = file.Save(filepath)
	if err != nil {
		return "", err
	}
	return filepath, nil
}
