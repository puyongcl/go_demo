package model

import (
	"gorm"
	"time"
)

/* data model*/

// demo_order
type Order struct {
	Id       uint      `json:"id" gorm:"primary_key"`
	OrderId  string    `json:"order_id" gorm:"type:varchar(64);not null;unique"`
	UserName string    `json:"user_name" gorm:"type:varchar(64)"`
	Amount   float64   `json:"amount" gorm:"type:float"`
	Status   string    `json:"status" gorm:"type:varchar(64)"`
	FileUrl  string    `json:"file_url" gorm:"type:varchar(256)"`
	CreateAt time.Time `json:"create_at" gorm:"column:create_at"`
}

func CheckTbl(db *gorm.DB) bool {
	return db.HasTable(&Order{})
}

func CreateTbl(db *gorm.DB) error {
	return db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Order{}).Error
}
