package handler

import (
	"github.com/gin-gonic/gin"
	"go_demo/model"
)

type BaseRSP struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"err_msg"`
	Info    string `json:"info"`
}

type QueryRsp struct {
	BaseRSP
	Data interface{} `json:"data"`
}

type QueryListRsp struct {
	BaseRSP
	Data []model.TOrder `json:"data"`
}

func SendSuccessRsp(c *gin.Context, data interface{}) {
	c.JSON(200, data)
}

func SendErrorRsp(c *gin.Context, errMsg string) {
	c.JSON(400, gin.H{"success": true, "err_msg": errMsg, "info": errMsg})
}

// 成功响应所有操作
func SendNormalRsp(c *gin.Context) {
	var rsp BaseRSP
	rsp.ErrMsg = ""
	rsp.Success = true
	rsp.Info = "success"
	SendSuccessRsp(c, rsp)
}
