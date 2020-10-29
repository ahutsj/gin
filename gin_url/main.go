package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取请求的path（URL）参数，返回的都是字符串类型、
//注意url的匹配不要冲突
func main() {
	r := gin.Default()
	//获取路径参数
	r.GET("/user/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK,gin.H{
			"name":name,
			"age":age,
		})
	})

	r.GET("/blogs/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK,gin.H{
			"year":year,
			"month":month,
		})
	})

	r.Run(":9090")
}
