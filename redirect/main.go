package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		/*c.JSON(http.StatusOK,gin.H{
			"status":"OK",
		})*/
		//跳转到 dogedoge.com
		c.Redirect(http.StatusMovedPermanently,"https://www.dogedoge.com")
	})

	//路由重定向
	r.GET("/a", func(c *gin.Context) {
		//跳转到 /b 对应的路由函数
		c.Request.URL.Path = "/b" //把请求的URI修改
		r.HandleContext(c) //继续后续处理
	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"message":"here is b", //地址还是 /a
		})
	})

	r.Run(":9090")
}
