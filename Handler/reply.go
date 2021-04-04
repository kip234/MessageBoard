package Handler

import (
	"MessageBoard/Model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Reply(db interface{}) gin.HandlerFunc {
	return func(c *gin.Context){
		var reply Model.Message
		err := c.ShouldBind(&reply)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"method":  "POST",
				"routing": "reply",
				"Error":   err.Error(),
			})
			return
		} else {
			//获取用户ID
			Uid,isExist:=c.Get("Uid")
			if !isExist{//不存在
				c.Redirect(http.StatusMovedPermanently,"/login")
				return
			}
			reply.Uid=Uid.(int)
			reply.Save(db)
			tmp:=Model.GetContent(db)
			c.JSON(http.StatusOK, &tmp)
		}
	}
}
