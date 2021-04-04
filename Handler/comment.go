package Handler

import (
	"MessageBoard/Model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Comment(db interface{}) gin.HandlerFunc {
	return func(c *gin.Context){
		var comment Model.Message
		//绑定参数
		if err := c.ShouldBind(&comment);err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"method":  "POST",
				"routing": "comment",
				"Error":   err.Error(),
			})
			return
		}
		//获取用户ID
		Uid,isExist:=c.Get("Uid")
		if !isExist{//不存在
			c.Redirect(http.StatusMovedPermanently,"/login")
			return
		}
		comment.Uid=Uid.(int)
		//回馈数据
		comment.Save(db)
		tmp := Model.GetContent(db)
		c.JSON(http.StatusOK,&tmp)
	}
}
