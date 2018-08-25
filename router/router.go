package router

import (
	"github.com/gin-gonic/gin"
	"go_demo/handler"
)

func RunRouter() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/orders", handler.AddOrderHandler)
		v1.POST("/orders/:order_id", handler.UpdateOrderHandler)
		v1.GET("/orders/:order_id", handler.GetOrderHandler)
		v1.GET("/orders/:user_name", handler.ListOrderByUserNameHandler)
		v1.POST("/orders/:order_id", handler.UploadFileHandler)
		v1.GET("orders/:order_id", handler.DownloadFileHandler)
	}
	router.Run(":9527")
}
