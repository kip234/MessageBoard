package Handler

import (
	"MessageBoard/services/DC"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/rpc"
	"strconv"
)

const AvailableLimit = 60*10//登录有效时限(秒)

//用户登录，成功后设置cookie
func Login(DCClient *rpc.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user DC.User
		err:=c.ShouldBind(&user)
		if err!=nil{//绑定失败
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"routing": "login",
				"Error":err.Error(),
			})
			return
		}
		tmp:= DC.User{}
		err=DCClient.Call("User.Load",user.Uid,&tmp)
		if err!=nil{//查找失败
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"routing": "login",
				"Error":err.Error(),
			})
			return
		}
		if tmp.Pwd!=user.Pwd{//密码对不上
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"routing": "login",
				"Error":"password wrong !",
				"tmp":tmp.Pwd,
				"user":user.Pwd,
			})
			return
		}
		//设置cookie
		c.SetCookie("Uid",strconv.Itoa(user.Uid),AvailableLimit,"/","localhost",false,true)
	}
}