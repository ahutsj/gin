package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//form表单
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html","./index.html")
	r.GET("/login", func(c *gin.Context) { //一次请求对应一个相应
		c.HTML(http.StatusOK,"login.html",nil)
	})

	//  /login post请求
	r.POST("/login", func(c *gin.Context) {
		//获取form表单提交的数据
		/*username := c.PostForm("username")
		password := c.PostForm("password")*/

		username := c.DefaultPostForm("username","somebody")
		password := c.DefaultPostForm("password","***")

		c.HTML(http.StatusOK,"index.html",gin.H{
			"Name":username,
			"Password":password,
		})
	})

	r.Run(":9090")

}
