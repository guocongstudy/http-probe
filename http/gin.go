package http

import (
	"github.com/gin-gonic/gin"
	"http-probe/config"
	"net/http"
)

func StartGin(c *config.Config) {
	//初始化gin实例
	r := gin.Default()
	//把路由绑定起来
	Routes(r)

	r.Run(c.HttpListenAddr)

}

//添加路由的函数
func Routes(r *gin.Engine) {
	api := r.Group("/api")
	api.GET("/probe/http", HttpProbe)
	api.GET("/v1", func(c *gin.Context) {
		c.String(http.StatusOK, "你好我是probe探针！")
	})

}
