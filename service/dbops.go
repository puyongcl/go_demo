package service

import (
	"errors"
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
	return dbConn.Create(rec).Error
}

func updateOrder(rec *model.Order) error {
	return dbConn.Model(&model.Order{}).Where("order_id = ?",
		rec.OrderId).Updates(map[string]interface{}{"amount": rec.Amount, "status": rec.Status,
		"file_url": rec.FileUrl}).Error
}

func getOrder(rec *model.Order) error {
	return dbConn.Where("order_id = ?", rec.OrderId).First(rec).Error
}

func getOrderListByUserName(key string, rec *[]model.Order) error {
	arg := fmt.Sprintf("%%%s%%", key)
	return dbConn.Where("user_name LIKE ?", arg).Order("create_at").Order("amount").Find(rec).Error
}

func updateOrderFileURL(rec *model.Order) error {
	return dbConn.Model(&model.Order{}).Where("order_id = ?",
		rec.OrderId).Update("file_url", rec.FileUrl).Error
}

func getOrderList(rec *[]model.Order) error {
	return dbConn.Find(rec).Error
}

func transAmount(from string, to string, amount float64) error {
	if from == "" || to == "" || amount == 0.0 {
		return errors.New("invalid args")
	}

	tx := dbConn.Begin()
	defer tx.Rollback()
	// query
	var recFrom model.Order
	err := tx.Where("order_id = ?", from).First(&recFrom).Error
	if err != nil {
		return err
	}
	if recFrom.Amount < amount {
		return errors.New("amount not enough")
	}
	var recTo model.Order
	err = tx.Where("order_id = ?", to).First(&recTo).Error
	if err != nil {
		return err
	}

	// update
	recFrom.Amount -= amount
	err = tx.Model(&model.Order{}).Where("order_id = ?", from).Update("amount", recFrom.Amount).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	recTo.Amount += amount
	err = tx.Model(&model.Order{}).Where("order_id = ?", to).Update("amount", recTo.Amount).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func getOrderPageListByUserName(username string, rec *[]model.Order, pageNo uint, size uint) (recordCnt uint, pageCnt uint, err error) {
	if pageNo == 0 {
		pageNo = 1
	}
	err = dbConn.Model(&model.Order{}).Where("user_name = ?",
		username).Count(&recordCnt).Limit(size).Offset((pageNo - 1) * size).Find(rec).Error
	pageCnt = (recordCnt + size - 1) / size
	return
}
