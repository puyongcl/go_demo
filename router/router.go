package router

import (
	"go_demo/handler"
	"gopkg.in/gin-gonic/gin.v1"
)

func RunRouter() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/orders", handler.AddOrderHandler)
		v1.POST("/orders/update/:order_id", handler.UpdateOrderHandler)
		v1.GET("/orders/get/:order_id", handler.GetOrderHandler)
		v1.GET("/orders/list/:name_key", handler.ListOrderByUserNameHandler)
		v1.POST("/orders/upload/:order_id", handler.UploadFileHandler)
		v1.GET("/orders/download/:order_id", handler.DownloadFileHandler)
		v1.POST("orders/list_page/", handler.ListOrderByUserNamePageHandler)
	}
	// TODO: 1. POST json实现一个接口：条件查询订单（分页，条件为：姓名与编号）
	// TODO: 2. 使用TCP库net/tcp，实现使用浏览器的GET请求-服务器
	// 3. 如果struct需要初始化，直接在声明时。
	// TODO: 4. route直接写参与body写参要理解区别。http get && post(json, form-data, x-www-form-urlencode)
	// 5. gin读取文件名
	// TODO: 6. gin Param与Query的函数区别 xx?k1=v1&k2=v2
	/*
		// Param returns the value of the URL param.
		// It is a shortcut for c.Params.ByName(key)
		//		router.GET("/user/:id", func(c *gin.Context) {
		//			// a GET request to /user/john
		//			id := c.Param("id") // id == "john"
		//		})

		// Query returns the keyed url query value if it exists,
		// othewise it returns an empty string `("")`.
		// It is shortcut for `c.Request.URL.Query().Get(key)`
		// 		GET /path?id=1234&name=Manu&value=
		// 		c.Query("id") == "1234"
		// 		c.Query("name") == "Manu"
		// 		c.Query("value") == ""
		// 		c.Query("wtf") == ""
	*/
	router.Run(":8080")
}
