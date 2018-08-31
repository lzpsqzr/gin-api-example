package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lzpsqzr/gin-api-example/routers/api/v1"
	"io"
	"os"
)

func Initrouters() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	// Use the following code if you need to write the logs only to file.
	// gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/alerts", v1.GetTags)
		//新建标签
		apiv1.POST("/alerts", v1.PostTag)
		//更新指定标签
		apiv1.PUT("/alerts/:id", v1.PutTag)
		//删除指定标签
		apiv1.DELETE("/alerts/:id", v1.DeleteTag)
	}

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pongpong")
	})

	return r
}
