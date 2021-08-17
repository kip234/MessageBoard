package Middleware

import (
	"MessageBoard/services/DC"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/rpc"
	"strconv"
)

//读取cookie
func IsLogin(DCClient *rpc.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user DC.User
		var ok bool
		Uid,err:=c.Cookie("Uid")
		if err!=nil{//没有登录
			c.Redirect(http.StatusMovedPermanently,"/login")
			c.Abort()
			return
		}
		user.Uid,err=strconv.Atoi(Uid)
		if err!=nil{//绑定失败
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"routing": "isLogin",
				"Uid":Uid,
				"Error":err.Error(),
			})
			c.Abort()
			return
		}else if err=DCClient.Call("User.IsExist",user,&ok);err!=nil{//用户不存在
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"routing": "isLogin",
				"Error":"unknown user !",
				"errorinfo":err.Error(),
			})
			c.Abort()
			return
		}
		//保存信息供后续使用
		c.Set("Uid",user.Uid)
	}
}