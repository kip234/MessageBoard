package Routers

import (
	"MessageBoard/Handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BuildRouter(db *gorm.DB) *gin.Engine {
	server:=gin.Default()

	group:=server.Group("/")
	{
		group.POST("/publish", Handler.Publish(db))
		group.POST("/comment", Handler.Comment(db))
		group.POST("/like", Handler.Like(db))
		group.POST("/reply", Handler.Reply(db))
	}

	//暂时没弄
	//server.POST("/register",Handler.Register(db))
	server.GET("/",Handler.Home(db))

	return server

}
