package Handler

import (
	"MessageBoard/services/DC"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/rpc"
)

func Home(DCClient *rpc.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var message []DC.Message
		err:=DCClient.Call("Message.GetContent",DC.Message{},&message)
		//回馈数据
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"Error:":err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,&message)
	}
}
