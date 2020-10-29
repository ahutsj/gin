package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})
	//上传的文件的大小限制
	//r.MaxMultipartMemory = 8 << 20 //8Mib
	r.POST("/upload", func(c *gin.Context) {
		//1.从请求中读取文件
		f,err := c.FormFile("f1") //从请求中获取携带的参数是一样的
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else {
			//2.将读取到的文件保存在本地（服务器本地）
			//dst := fmt.Sprintf("./%s",f.Filename)
			dst := path.Join("./",f.Filename)
			_ = c.SaveUploadedFile(f,dst)
			c.JSON(http.StatusOK,gin.H{
				"status":"OK",
			})
		}
	})

	r.Run(":9090")
}
