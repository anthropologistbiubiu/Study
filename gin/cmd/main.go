package main

import (
	"gin/upload"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.POST("/upload", upload.UploadFileControl)
	r.GET("/download", upload.DownLoadFile)
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")

}
