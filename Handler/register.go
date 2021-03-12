//暂时没弄
package Handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"MessageBoard/Model"
)

func Register(db *gorm.DB) gin.HandlerFunc {
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
			if tmp:=db.Create(&user);tmp.Error!=nil{//存入数据库出错
				c.JSON(http.StatusBadRequest,gin.H{
					"method":  "POST",
					"routing": "register",
					"Error":tmp.Error.Error(),
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
