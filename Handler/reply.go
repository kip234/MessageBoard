package Handler

import (
	"MessageBoard/services/DC"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/rpc"
)

func Reply(DCClient *rpc.Client) gin.HandlerFunc {
	return func(c *gin.Context){
		var reply DC.Message
		err := c.ShouldBind(&reply)
		if err != nil {//绑定失败
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
			DCClient.Call("Message.Save",reply,&DC.Message{})
			//reply.Save(db)
			//反馈信息
			var tmp []DC.Message
			DCClient.Call("Message.GetContent",&DC.Message{},&tmp)
			//tmp:= DC.GetContent(db)
			c.JSON(http.StatusOK, &tmp)
		}
	}
}
