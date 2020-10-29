package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//GET请求 URL ?后面的事情
	//key = value 的格式
	//多个 key value 使用 & 连接
	//examp : /web?query=程潇&age=18
	r.GET("/web", func(c *gin.Context){
		//获取浏览器那边发请求携带的 query string 参数
		//方法一
		name := c.Query("query") //通过Query获取请求中携带的query string参数
		age := c.Query("age")

		//方法二
		/*name := c.DefaultQuery("query","somebody") //取不到就用指定的默认值*/

		//方法三
		/*name,OK := c.GetQuery("query" ) //取到就返回(value,true)，取不到的话返回("",false)
		if !OK {
			//取不到
			name = "somebody"
		}*/
		c.JSON(http.StatusOK,gin.H{
			"name":name,
			"age":age,
		})
	})

	r.Run(":9090")
}
