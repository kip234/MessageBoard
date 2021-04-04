//暂时没弄
package Handler

import (
	"MessageBoard/Model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//注册成狗后会返回UID
func Register(db interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Model.User
		err:=c.ShouldBind(&user)
		if err!=nil {//绑定出错
			c.JSON(http.StatusBadRequest,gin.H{
				"method":  "POST",
				"routing": "register",
				"Error":err.Error(),
			})
		}else {
			if err:=user.Save(db);err!=nil{//存入数据库出错
				c.JSON(http.StatusBadRequest,gin.H{
					"method":  "POST",
					"routing": "register",
					"Error":err.Error(),
				})
			}else{
				c.JSON(http.StatusOK,gin.H{
					"method":  "POST",
					"routing": "register",
					"AddUser":user,
				})
			}
		}
	}
}
