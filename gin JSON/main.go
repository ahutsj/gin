package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	r.GET("/JSON",func(c *gin.Context){
		//方法一：使用map
		data := map[string]interface{}{
			"name":"小王子",
			"age":18,
			"message":"hello earth!",
		}

		/*data := gin.H{
			"name":"小王子",
			"message":"hello earth",
			"age":18,
		}*/

		c.JSON(http.StatusOK,data)
	})

	//方法二：使用结构体 灵活使用tag来做定制化操作
	type msg struct{
		Name string `json:"name"`
		Age int
		Message string
	}

	r.GET("/another_JSON",func(c *gin.Context){
		data := msg{
			Name: "小王子",
			Age: 18,
			Message: "hello earth!!",
		}
		c.JSON(http.StatusOK,data) //JSON的序列化，Go语言中首字母小写表示不可导出，如果必须要返回小写，使用 tag
		/*type msg struct{
			Name: string `json:"name"` //以此类推 使用 tag 进行标记
			Age: int



		
			Message: string
		}*/
	})

	r.Run(":9090")
}
