package main

import "github.com/gin-gonic/gin"

func sayhello(c *gin.Context) {
	c.JSON(200,gin.H{
		"message":"hello gin",
	})
}

func main() {

	//返回默认的路由引擎
	r := gin.Default()

	//指定用户使用GET请求访问
	r.GET("/hello",sayhello)

	//启动服务
	r.Run()
}
