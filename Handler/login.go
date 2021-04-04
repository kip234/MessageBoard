package Handler

import (
	"MessageBoard/Model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const AvailableLimit = 60*10//登录有效时限(秒)

func Login(db interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Model.User
		err:=c.ShouldBind(&user)
		if err!=nil{//绑定失败
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"routing": "register",
				"Error":err.Error(),
			})
			return
		}
		tmp:=Model.User{}
		if err:=tmp.Load(db,user.Uid);err!=nil{//查找失败
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"routing": "register",
				"Error":err.Error(),
			})
			return
		}
		if tmp.Pwd!=user.Pwd{//密码对不上
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"routing": "register",
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