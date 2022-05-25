package router

import (
	"bluebell/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()

	// 引入日志中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 每两秒钟添加一个令牌  全局限流
	//r.Use(logger.GinLogger(), logger.GinRecovery(true),middlewares.RateLimitMiddleware(2*time.Second , 1))

	// 接口文档：http://localhost:8083/swagger/index.html#/
	//r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	//v1 := r.Group("/api/v1")
	r.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
