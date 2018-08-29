package handler

import (
	"go_demo/model"
	"gopkg.in/gin-gonic/gin.v1"
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
	Data []model.Order `json:"data"`
}

type QueryListPageRsp struct {
	BaseRSP
	Data        []model.Order `json:"data"`
	RecordCount uint          `json:"record_count"`
	PageCount   uint          `json:"page_count"`
}

func SendSuccessRsp(c *gin.Context, data interface{}) {
	c.JSON(200, data)
}

func SendErrorRsp(c *gin.Context, errMsg string) {
	c.JSON(400, gin.H{"success": true, "err_msg": errMsg, "info": errMsg})
}

// 成功响应所有操作
func SendNormalRsp(c *gin.Context) {
	SendSuccessRsp(c, BaseRSP{
		Success: true,
		Info:    "success",
	})
}
