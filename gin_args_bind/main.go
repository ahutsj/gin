package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username" json:"username"`  //tag的作用是查找时可以一一对应
	Password string `form:"password" json:"password"`
}
func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./index.html")
	r.GET("/user", func(c *gin.Context) {
		/*username := c.Query("username")
		password := c.Query("password")
		u := UserInfo{
			username: username,
			password: password,
		}*/
		var u UserInfo //声明一个结构体
		err := c.ShouldBind(&u) //遍历结构体进行字段匹配，必须传一个指针进去，另外结构体想要被自动匹配，就需要反射机制，即结构体首字母需要大写
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}else {
			fmt.Printf("%#v\n",u)
			c.JSON(http.StatusOK,gin.H{
				"status":"ok",
			})
		}
		/*fmt.Printf("%#v\n",u)*/
		/*c.JSON(http.StatusOK,gin.H{
			"message":"ok",
		})*/
	})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})

	r.POST("/form", func(c *gin.Context) {
		var u UserInfo //声明一个结构体
		err := c.ShouldBind(&u) //遍历结构体进行字段匹配，必须传一个指针进去，另外结构体想要被自动匹配，就需要反射机制，即结构体首字母需要大写
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}else {
			fmt.Printf("%#v\n",u)
			c.JSON(http.StatusOK,gin.H{
				"status":"ok",
			})
		}
	})

	r.POST("/json", func(c *gin.Context) {
		var u UserInfo //声明一个结构体
		err := c.ShouldBind(&u) //遍历结构体进行字段匹配，必须传一个指针进去，另外结构体想要被自动匹配，就需要反射机制，即结构体首字母需要大写
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}else {
			fmt.Printf("%#v\n",u)
			c.JSON(http.StatusOK,gin.H{
				"status":"ok",
			})
		}
	})

	r.Run(":9090")
}
