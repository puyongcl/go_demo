package model

import ("github.com/gin-gonic/gin")

type BaseRSP struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"err_msg"`
	Info    string `json:"info"`
}

func SendSuccessRsp(c *gin.Context, data interface{}) {
	c.JSON(200, data)
}

func SendErrorRsp(c *gin.Context, errMsg string) {
	c.JSON(200, gin.H{"success": true, "err_msg": errMsg, "info": errMsg})
}

type QueryRsp struct {
	BaseRSP
	Data interface{} `json:"data"`
}

type QueryListRsp struct {
	BaseRSP
	Data []TOrder `json:"data"`
}
