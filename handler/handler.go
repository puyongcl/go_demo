package handler

import (
	"fmt"
	"go_demo/common/util"
	"go_demo/model"
	"go_demo/service"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"strconv"
)

// 创建 demo_order
func AddOrderHandler(c *gin.Context) {
	fmt.Println("--------1---------")
	username := c.PostForm("username")
	status := c.PostForm("status")
	file, header, err := c.Request.FormFile("file")
	amount, err := strconv.ParseFloat(c.PostForm("amount"), 64)
	fmt.Println("1:", username, "2:", status, "3:", amount)
	if err != nil {
		err1 := "parse amount error：" + err.Error()
		SendErrorRsp(c, err1)
		return
	}

	// gen order id
	orderid, err := util.NewOrderID()
	if err != nil {
		err1 := "gen order id error：" + err.Error()
		SendErrorRsp(c, err1)
		return
	}

	if err = service.AddNewOrder(orderid, username, amount, status, file, header.Filename, service.FileDir); err != nil {
		err1 := "add new order error：" + err.Error()
		SendErrorRsp(c, err1)
		return
	}

	SendNormalRsp(c)
	return
}

// 更新order（amount、status、file_url）
func UpdateOrderHandler(c *gin.Context) {
	orderid := c.Param("order_id")
	status := c.PostForm("status")
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		err1 := "request file error：" + err.Error()
		SendErrorRsp(c, err1)
	}
	amount, err := strconv.ParseFloat(c.PostForm("amount"), 64)
	if err != nil {
		err1 := "parse amount error：" + err.Error()
		SendErrorRsp(c, err1)
		return
	}

	if err = service.UpdateOrder(orderid, amount, status, file, header.Filename, service.FileDir); err != nil {
		err1 := "add new order error：" + err.Error()
		SendErrorRsp(c, err1)
		return
	}

	SendNormalRsp(c)
}

// 获取order详情
func GetOrderHandler(c *gin.Context) {
	orderid := c.Param("order_id")
	order := model.Order{OrderId: orderid}
	if err := service.GetOrder(&order); err != nil {
		err1 := "get order error：" + err.Error()
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
	usernameKey := c.Param("name_key")
	var order []model.Order
	if err := service.GetOrderListByUserName(usernameKey, &order); err != nil {
		err1 := "get order list by username error：" + err.Error()
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
	orderid := c.Param("order_id")
	log.Println(orderid)
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		err1 := "request file error：" + err.Error()
		SendErrorRsp(c, err1)
		return
	}

	// upload file
	err = service.UploadFile(file, header.Filename, orderid, service.FileDir)
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
	orderid := c.Param("order_id")
	fmt.Println("------------------", orderid)
	fileURL, err := service.GetOrderFileURL(orderid)
	if err != nil {
		err1 := "get file path error：" + err.Error()
		SendErrorRsp(c, err1)
		return
	}
	c.File(fileURL)

	//SendNormalRsp(c) 返回文件时不再需要
}

// 根据姓名查找所有订单，需要分页
func ListOrderByUserNamePageHandler(c *gin.Context) {
	//res, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	err1 := "read body error：" + err.Error()
	//	SendErrorRsp(c, err1)
	//}
	//
	var ubody model.OrderPageListReq
	//if err = json.Unmarshal(res, ubody); err != nil {
	//	err1 := "unmarshal json body error：" + err.Error()
	//	SendErrorRsp(c, err1)
	//}
	if err := c.BindJSON(&ubody); err != nil {
		err1 := "parse json error：" + err.Error()
		SendErrorRsp(c, err1)
	}

	var order []model.Order
	recordCnt, pageCnt, err := service.GetOrderLIstByUserNamePage(ubody.Username, &order, ubody.PageNo, ubody.Size)
	if err != nil {
		err1 := "get order list by username error：" + err.Error()
		SendErrorRsp(c, err1)
		return
	}

	var rsp QueryListPageRsp
	rsp.ErrMsg = ""
	rsp.Success = true
	rsp.Info = "success"
	rsp.Data = order
	rsp.PageCount = pageCnt
	rsp.RecordCount = recordCnt
	SendSuccessRsp(c, rsp)
}
