package Handler

import (
	"MessageBoard/services/DC"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/rpc"
)

func Comment(DCClient *rpc.Client) gin.HandlerFunc {
	return func(c *gin.Context){
		var comment DC.Message
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
		DCClient.Call("Message.Save",comment,&DC.Message{})
		//comment.Save(db)
		var tmp []DC.Message
		//tmp := Message.GetContent(db)
		DCClient.Call("Message.GetContent",DC.Message{},&tmp)
		c.JSON(http.StatusOK,&tmp)
	}
}
