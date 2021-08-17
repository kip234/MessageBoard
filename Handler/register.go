//暂时没弄
package Handler

import (
	"MessageBoard/services/DC"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/rpc"
)

//注册成狗后会返回UID
func Register(DCClient *rpc.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user DC.User
		err:=c.ShouldBind(&user)
		if err!=nil {//绑定出错
			c.JSON(http.StatusBadRequest,gin.H{
				"method":  "POST",
				"routing": "register",
				"Error":err.Error(),
			})
		}else {
			re:=DC.User{}
			err:=DCClient.Call("User.Save",user,&re)
			if err!=nil{//存入数据库出错
				c.JSON(http.StatusBadRequest,gin.H{
					"method":  "POST",
					"routing": "register",
					"Error":err.Error(),
				})
			}else{
				c.JSON(http.StatusOK,gin.H{
					"method":  "POST",
					"routing": "register",
					"AddUser":re,
				})
			}
		}
	}
}
