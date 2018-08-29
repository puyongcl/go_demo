package service

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"go_demo/common/util"
	"go_demo/model"
	"log"
	"mime/multipart"
	"os"
	"time"
)

func init() {
	// 判断上传文件目录是否存在，不存在则创建
	_, err := os.Stat(FileDir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(FileDir, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	}
}

func AddNewOrder(orderid string, username string, amount float64, status string, file multipart.File, filename string, fileDir string) (err error) {
	//if orderid == "" || username == "" || amount == 0.0 || status == "" {
	//	err = errors.New("username,amount,status can't be empty")
	//	log.Println(err.Error())
	//	return
	//}

	// save file
	fileURL, err := util.SaveFile(file, filename, fileDir)
	if err != nil {
		err = errors.New("save file error：" + err.Error())
		log.Println(err.Error())
		return
	}

	// insert
	err = insertNewOrderRecord(&model.Order{OrderId: orderid, UserName: username, Amount: amount,
		Status: status, FileUrl: fileURL, CreateAt: time.Now()})
	if err != nil {
		err = errors.New("insert a new record error:" + err.Error())
		log.Println(err.Error())
	}
	return
}

func UpdateOrder(orderid string, amount float64, status string, file multipart.File, filename string, fileDir string) (err error) {
	//if orderid == "" || amount == 0.0 || status == "" {
	//	err = errors.New("username,amount,status can not be empty")
	//	log.Println(err.Error())
	//	return
	//}

	// save file
	fileURL, err := util.SaveFile(file, filename, fileDir)
	if err != nil {
		err = errors.New("save file error：" + err.Error())
		log.Println(err.Error())
		return
	}

	// get old record and del old file
	orderold := model.Order{OrderId: orderid}
	err = getOrder(&orderold)
	if err != nil {
		err = errors.New("get old record error:" + err.Error())
		log.Println(err.Error())
		return
	}
	if orderold.FileUrl != "" {
		err = os.Remove(orderold.FileUrl)
		if err != nil {
			err = errors.New("remove old file error:" + err.Error())
			log.Println(err.Error())
			return
		}
	}

	// update DB
	err = updateOrder(&model.Order{OrderId: orderid, Amount: amount, Status: status, FileUrl: fileURL})
	if err != nil {
		err = errors.New("update order error:" + err.Error())
		log.Println(err.Error())
	}
	return
}

func GetOrder(rec *model.Order) error {
	return getOrder(rec)
}

func GetOrderLIstByUserNamePage(username string, rec *[]model.Order, pageNo uint, size uint) (recordCnt uint, pageCnt uint, err error) {
	//if username == "" {
	//	err = errors.New("username can not be empty")
	//	return
	//}
	recordCnt, pageCnt, err = getOrderPageListByUserName(username, rec, pageNo, size)
	if err != nil {
		err = errors.New("get order list by username and paging error:" + err.Error())
		log.Println(err.Error())
	}
	return
}

func GetOrderListByUserName(key string, rec *[]model.Order) (err error) {
	//if key == "" {
	//	err = errors.New("username can not be empty")
	//	log.Println(err.Error())
	//	return
	//}
	err = getOrderListByUserName(key, rec)
	if err != nil {
		err = errors.New("get order list error:" + err.Error())
		log.Println(err.Error())
	}
	return
}

func UpdateOrderFileURL(orderid string, fileURL string) (err error) {
	if orderid == "" || fileURL == "" {
		err = errors.New("order id and fileURL can not be empty")
		log.Println(err.Error())
		return
	}

	return updateOrderFileURL(&model.Order{OrderId: orderid, FileUrl: fileURL})
}

func GetOrderFileURL(orderid string) (fileURL string, err error) {
	if orderid == "" {
		err = errors.New("order id can not be empty")
		log.Println(err.Error())
		return
	}

	order := model.Order{OrderId: orderid}
	err = getOrder(&order)
	if err != nil {
		err = errors.New("get order file URL error:" + err.Error())
		log.Println(err.Error())
	}
	fileURL = order.FileUrl
	return
}

func UploadFile(file multipart.File, filename string, orderid string, filedir string) (err error) {
	// save file
	fileURL, err := util.SaveFile(file, filename, filedir)
	if err != nil {
		err = errors.New("save file error：" + err.Error())
		log.Println(err.Error())
		return
	}

	// update DB file_url
	if err = UpdateOrderFileURL(orderid, fileURL); err != nil {
		err = errors.New("update file URL error：" + err.Error())
		log.Println(err.Error())
		return
	}
	return
}

func ExportOrderListWithExcel() (excelFileURL string, err error) {
	// get all order record
	var rec []model.Order
	err = getOrderList(&rec)
	if err != nil {
		err = errors.New("get all order record error:" + err.Error())
		log.Println(err.Error())
		return
	}

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		err = errors.New("add sheet error:" + err.Error())
		log.Println(err.Error())
		return
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

	excelFileURL, err = util.NewUUID()
	if err != nil {
		err = errors.New("gen file name error:" + err.Error())
		log.Println(err.Error())
		return
	}
	excelFileURL += ".xlsx"
	excelFileURL = FileDir + excelFileURL
	err = file.Save(excelFileURL)
	if err != nil {
		err = errors.New("save excel file error:" + err.Error())
		log.Println(err.Error())
		return
	}
	return
}
