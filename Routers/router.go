//建立路由结构
//为了便于测试几乎给每条路由都加了get方法
package Routers

import (
	"MessageBoard/Handler"
	"MessageBoard/Middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BuildRouter(db interface{}) *gin.Engine {
	server:=gin.Default()

	group:=server.Group("/",Middleware.IsLogin(db))
	{
		group.POST("/publish", Handler.Publish(db))
		group.GET("/publish", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "publish",
			})
		})

		group.POST("/comment", Handler.Comment(db))
		group.GET("/comment", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "comment",
			})
		})

		group.POST("/like", Handler.Like(db))
		group.GET("/like", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "like",
			})
		})

		group.POST("/reply", Handler.Reply(db))
		group.GET("/reply", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "reply",
			})
		})
	}

	server.POST("/register",Handler.Register(db))
	server.GET("/register",func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":  "GET",
			"routing": "register",
		})
	})

	server.POST("/login",Handler.Login(db))
	server.GET("/login",func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":  "GET",
			"routing": "login",
		})
	})

	server.GET("/",Handler.Home(db))

	return server

}
