package service

import (
	"fmt"
	"go_demo/db"
	"go_demo/model"
	"gorm"
)

var dbConn *gorm.DB

func init() {
	dbConn = db.GetDB()
	if dbConn == nil {
		panic("can't get db conn!")
	}
}

func insertNewOrderRecord(rec *model.Order) error {
	return dbConn.Create(&rec).Error
}

func updateOrder(rec *model.Order) error {
	return dbConn.Model(&rec).Where("order_id = ?", rec.OrderId).Updates(map[string]interface{}{"amount": rec.Amount, "status": rec.Status, "file_url": rec.FileUrl}).Error
}

func getOrder(rec *model.Order) error {
	return dbConn.Where("order_id = ?", rec.OrderId).First(rec).Error
}

func getOrderListByUserName(key string, rec []model.Order) error {
	arg := fmt.Sprintf("%%%s%%", key)
	return dbConn.Where("user_name LIKE ?", arg).Order("create_at").Order("amount").Find(&rec).Error
}

func updateOrderFileURL(rec *model.Order) error {
	return dbConn.Model(&rec).Where("order_id = ?", rec.OrderId).Update("file_url", rec.FileUrl).Error
}
