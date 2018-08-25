package handler

import (
	"github.com/gin-gonic/gin"
	"go_demo/common/util"
	"go_demo/model"
	"go_demo/service"
	"log"
	"strconv"
)

// 创建 demo_order
func AddOrderHandler(c *gin.Context) {
	username := c.PostForm("username")
	status := c.PostForm("status")
	amount, err := strconv.ParseFloat(c.PostForm("amount"), 64)
	if err != nil {
		err1 := "parse amount error：" + err.Error()
		log.Println(err1)
		SendErrorRsp(c, err1)
		return
	}

	// gen order id
	orderid, err := util.NewOrderID()
	if err != nil {
		err1 := "gen order id error：" + err.Error()
		log.Println(err1)
		SendErrorRsp(c, err1)
		return
	}

	if err = service.AddNewOrder(orderid, username, amount, status, ""); err != nil {
		err1 := "add new order error：" + err.Error()
		log.Println(err1)
		SendErrorRsp(c, err1)
		return
	}

	SendNormalRsp(c)
}

// 更新order（amount、status、file_url）
func UpdateOrderHandler(c *gin.Context) {
	orderid := c.PostForm("order_id")
	status := c.PostForm("status")
	fileURL := c.PostForm("file_url")
	amount, err := strconv.ParseFloat(c.PostForm("amount"), 64)
	if err != nil {
		err1 := "parse amount error：" + err.Error()
		log.Println(err1)
		SendErrorRsp(c, err1)
		return
	}

	if err = service.UpdateOrder(orderid, amount, status, fileURL); err != nil {
		err1 := "add new order error：" + err.Error()
		log.Println(err1)
		SendErrorRsp(c, err1)
		return
	}

	SendNormalRsp(c)
}

// 获取order详情
func GetOrderHandler(c *gin.Context) {
	orderid := c.Query("order_id")
	order := model.Order{OrderId: orderid}
	if err := service.GetOrder(&order); err != nil {
		err1 := "get order error：" + err.Error()
		log.Println(err1)
		SendErrorRsp(c, err1)
		return
	}

	var rsp QueryRsp
	rsp.ErrMsg = ""
	rsp.Success = true
	rsp.Info = "success"
	rsp.Data = order
	SendSuccessRsp(c, rsp)
}

// 获取 demo_order 列表 （需要包含： 模糊查找、根据创建时间，金额排序）
func ListOrderByUserNameHandler(c *gin.Context) {
	usernameKey := c.Query("user_name_fuzzy_lookup_key")
	var order []model.Order
	if err := service.GetOrderListByUserName(usernameKey, order); err != nil {
		err1 := "get order list by username error：" + err.Error()
		log.Println(err1)
		SendErrorRsp(c, err1)
		return
	}

	var rsp QueryListRsp
	rsp.ErrMsg = ""
	rsp.Success = true
	rsp.Info = "success"
	rsp.Data = order
	SendSuccessRsp(c, rsp)
}

// 上传order中的file
func UploadFileHandler(c *gin.Context) {
	orderid := c.PostForm("order_id")

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		err1 := "request file error：" + err.Error()
		log.Println(err1)
		SendErrorRsp(c, err1)
		return
	}

	// upload file
	err = service.UploadFile(file, orderid, service.FileDir)
	if err != nil {
		err1 := "upload file error：" + err.Error()
		log.Println(err1)
		SendErrorRsp(c, err1)
		return
	}

	SendNormalRsp(c)
}

// 获取order中的file
func DownloadFileHandler(c *gin.Context) {
	orderid := c.Query("order_id")
	filepath, err := service.GetOrderFileURL(orderid)
	if err != nil {
		err1 := "get file path error：" + err.Error()
		log.Println(err1)
		SendErrorRsp(c, err1)
		return
	}
	c.File(filepath)

	SendNormalRsp(c)
}
