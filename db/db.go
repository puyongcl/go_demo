package db

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"go_demo/db/config"
	"go_demo/model"
	"gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// init DB ,check if connect
	dbc := config.GetDBCfg()
	db, err = connDB(dbc)
	if err != nil {
		if err = createDB(dbc); err != nil {
			panic(err)
		}

		// reconnect db
		if db, err = connDB(dbc); err != nil {
			panic(err)
		}
	}

	// init tbl, check if exist
	if ok := model.CheckTbl(db); !ok {
		if err = model.CreateTbl(db); err != nil {
			panic(err)
		}
	}
}

func createDB(dbc *config.DBCfg) error {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbc.Username, dbc.Pwd, dbc.Addr, dbc.Port, "information_schema")
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		return errors.New("open fail：" + err.Error())
	}
	defer db.Close()

	sql := "create database if not exists " + dbc.DBName
	err = db.Exec(sql).Error
	if err != nil {
		return errors.New("create database fail：" + err.Error())
	}

	return nil
}

func connDB(dbc *config.DBCfg) (*gorm.DB, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbc.Username, dbc.Pwd, dbc.Addr, dbc.Port, dbc.DBName)
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetDB() *gorm.DB {
	return db
}

func clearTables() {
	db.Exec("truncate order")
}
