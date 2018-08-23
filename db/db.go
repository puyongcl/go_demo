package db

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)", dbc.Username, dbc.Pwd, dbc.Addr, dbc.Port)
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		return err
	}
	defer db.Close()

	sql := "create database " + dbc.DBName
	err = db.Exec(sql).Error
	if err != nil {
		return err
	}

	return nil
}

func connDB(dbc *config.DBCfg) (*gorm.DB, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
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
