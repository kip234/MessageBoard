package Handler

import (
	"MessageBoard/services/DC"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/rpc"
)

func Publish(DCClient *rpc.Client) gin.HandlerFunc {
	return func(c *gin.Context){
		var speech DC.Message
		err := c.ShouldBind(&speech)
		if err != nil {//绑定失败
			c.JSON(http.StatusOK, gin.H{
				"method":  "POST",
				"routing": "publish",
				"Error":   err.Error(),
			})
			return
		} else {
			//获取用户ID
			Uid,isExist:=c.Get("Uid")
			if !isExist{//不存在
				c.JSON(http.StatusOK,gin.H{
					"method":"GET",
					"router":"publish",
					"isExist":"isExist",
				})
				//c.Redirect(http.StatusMovedPermanently,"/login")
				return
			}
			speech.Uid=Uid.(int)
			tmp:=DCClient.Call("Message.Save",speech,&DC.Message{})
			if tmp!=nil{//保存失败
				c.JSON(http.StatusBadRequest, gin.H{
					"method":  "POST",
					"routing": "publish",
					"Error": tmp.Error(),
				})
				return
			}
		}
		var tmp []DC.Message
		DCClient.Call("Message.GetContent",DC.Message{},&tmp)
		c.JSON(http.StatusOK,&tmp)
	}
}