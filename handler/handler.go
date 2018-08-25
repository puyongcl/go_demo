package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"go_demo/model"
	"fmt"
	"log"
	"go_demo/service"
	"io/ioutil"
)

func AddOrderHandler(c *gin.Context) {
	username := c.PostForm("username")
	amount, err := strconv.ParseFloat(c.PostForm("amount"), 64)
	if err != nil {
		err1 := "parse amount error：" + err.Error()
		log.Println(err1)
		model.SendErrorRsp(c, err1)
		return
	}
	status := c.PostForm("status")

	if err = service.AddNewOrder(username, amount, status, ""); err != nil {
		err1 := "add new order error：" + err.Error()
		log.Println(err1)
		model.SendErrorRsp(c, err1)
		return
	}

	fmt.Println(amount, username)
	var rsp model.BaseRSP
	rsp.ErrMsg = ""
	rsp.Success = true
	rsp.Info = "success"
	model.SendSuccessRsp(c, rsp)
}

func UpdateOrderHandler(c *gin.Context) {
	orderid := c.PostForm("order_id")
	amount, err := strconv.ParseFloat(c.PostForm("amount"), 64)
	if err != nil {
		err1 := "parse amount error：" + err.Error()
		log.Println(err1)
		model.SendErrorRsp(c, err1)
		return
	}
	status := c.PostForm("status")
	fileURL := c.PostForm("file_url")

	if err = service.UpdateOrder(orderid, amount, status, fileURL); err != nil {
		err1 := "add new order error：" + err.Error()
		log.Println(err1)
		model.SendErrorRsp(c, err1)
		return
	}

	var rsp model.BaseRSP
	rsp.ErrMsg = ""
	rsp.Success = true
	rsp.Info = "success"
	model.SendSuccessRsp(c, rsp)
}

func GetOrderHandler(c *gin.Context) {
	orderid := c.Query("order_id")
	order := model.TOrder{OrderId: orderid}
	if err := service.GetOrder(&order); err != nil {
		err1 := "get order error：" + err.Error()
		log.Println(err1)
		model.SendErrorRsp(c, err1)
		return
	}

	var rsp model.QueryRsp
	rsp.ErrMsg = ""
	rsp.Success = true
	rsp.Info = "success"
	rsp.Data = order
	model.SendSuccessRsp(c, rsp)
}

// 获取 demo_order 列表 （需要包含： 模糊查找、根据创建时间，金额排序）
func ListOrderByUserNameHandler(c *gin.Context) {
	usernameKey := c.Query("user_name_fuzzy_lookup_key")
	var order []model.TOrder
	if err := service.GetOrderListByUserName(usernameKey, order); err != nil {
		err1 := "get order list by username error：" + err.Error()
		log.Println(err1)
		model.SendErrorRsp(c, err1)
		return
	}

	var rsp model.QueryListRsp
	rsp.ErrMsg = ""
	rsp.Success = true
	rsp.Info = "success"
	rsp.Data = order
	model.SendSuccessRsp(c, rsp)
}

func UploadFileHandler(c *gin.Context) {
	orderid := c.PostForm("order_id")

	file ,_,err:= c.Request.FormFile("file")
	if err != nil {
		err1 := "request file error：" + err.Error()
		log.Println(err1)
		model.SendErrorRsp(c, err1)
		return
	}

	// process file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		err1 := "read file error：" + err.Error()
		log.Println(err1)
		model.SendErrorRsp(c, err1)
		return
	}

	filename, err := service.NewUUID()
	if err != nil {
		err1 := "create file name error：" + err.Error()
		log.Println(err1)
		model.SendErrorRsp(c, err1)
		return
	}
	filepath := service.FileDir + filename
	err = ioutil.WriteFile(filepath, data, 0666)
	if err != nil {
		err1 := "write file error：" + err.Error()
		log.Println(err1)
		model.SendErrorRsp(c, err1)
	}

	// update DB file_url
	if err = service.UpdateOrderFileURL(orderid, filepath); err != nil {
		err1 := "update file path error：" + err.Error()
		log.Println(err1)
		model.SendErrorRsp(c, err1)
		return
	}

	var rsp model.BaseRSP
	rsp.ErrMsg = ""
	rsp.Success = true
	rsp.Info = "success"
	model.SendSuccessRsp(c, rsp)
}

func DownloadFileHandler(c *gin.Context) {
	orderid := c.Query("order_id")
	filepath := service.GetOrderFilePath(orderid)
	c.File("")
}