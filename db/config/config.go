package config

import (
	"encoding/json"
	"os"
)

type DBCfg struct {
	Addr     string `json:"db_addr"`
	Port     string `json:"db_port"`
	Username string `json:"db_user_name"`
	Pwd      string `json:"db_pwd"`
	DBName   string `json:"db_name"`
}

var dbconfig *DBCfg

func init() {
	//root := os.Getenv("GOPATH")
	//path := root + "/go_demo/bin/conf.json"
	path := "/home/qydev/workspace/go/src/go_demo/bin/conf.json"
	//dbconfig = &DBCfg{"localhost", 3306, "root", "root", "test"}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	dbconfig = &DBCfg{}

	err = decoder.Decode(dbconfig)
	if err != nil {
		panic(err)
	}
}

func GetDBCfg() *DBCfg {
	return dbconfig
}
