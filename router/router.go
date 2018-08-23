package router

import (
	"github.com/gin-gonic/gin"
	"go_demo/handler"
)


func RunRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/user/add_demo_order", handler.AddDemoOrder)
		v1.PUT("/user/update_demo_order", handler.UpdateDemoOrder)
		v1.GET("/user/get_demo_order", handler.GetDemoOrder)
		v1.GET("/user/get_demo_order_list", handler.ListDemoOrder)
	}
	router.Run(":9527")
}