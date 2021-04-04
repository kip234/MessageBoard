package Middleware

import (
	"MessageBoard/Model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func IsLogin(db interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Model.User
		Uid,err:=c.Cookie("Uid")
		if err!=nil{//没有登录
			/*c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"routing": "isLogin",
				"Error":err.Error(),
			})*/
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
			//c.Redirect(http.StatusMovedPermanently,"/login")
			c.Abort()
			return
		}else if !user.IsExist(db){//用户不存在
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"routing": "isLogin",
				"Error":"unknown user !",
			})
			//c.Redirect(http.StatusMovedPermanently,"/register")
			c.Abort()
			return
		}
		c.Set("Uid",user.Uid)
	}
}