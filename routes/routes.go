package routes

import (
	"net/http"
	"web_app/controller"
	"web_app/logger"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//注册业务路由
	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
