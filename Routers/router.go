//建立路由结构
//为了便于测试几乎给每条路由都加了get方法
package Routers

import (
	"MessageBoard/Handler"
	"MessageBoard/Middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/rpc"
)

func BuildRouter(DCClient *rpc.Client) *gin.Engine {
	server:=gin.Default()

	group:=server.Group("/",Middleware.IsLogin(DCClient))
	{
		group.POST("/publish", Handler.Publish(DCClient))
		group.GET("/publish", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "publish",
			})
		})

		group.POST("/comment", Handler.Comment(DCClient))
		group.GET("/comment", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "comment",
			})
		})

		group.POST("/like", Handler.Like(DCClient))
		group.GET("/like", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "like",
			})
		})

		group.POST("/reply", Handler.Reply(DCClient))
		group.GET("/reply", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "reply",
			})
		})
	}

	server.POST("/register",Handler.Register(DCClient))
	server.GET("/register",func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":  "GET",
			"routing": "register",
		})
	})

	server.POST("/login",Handler.Login(DCClient))
	server.GET("/login",func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":  "GET",
			"routing": "login",
		})
	})

	server.GET("/",Handler.Home(DCClient))

	return server

}
