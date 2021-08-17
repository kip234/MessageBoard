package Handler

import (
	"MessageBoard/services/DC"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/rpc"
	"strconv"
)

func Like(DCClient *rpc.Client) gin.HandlerFunc {
	return func(c *gin.Context){
		Pid,ok:=c.GetPostForm("Pid")
		//fmt.Println(Pid)
		if !ok{
			c.JSON(http.StatusOK, gin.H{
				"method":  "POST",
				"routing": "like",
				"Status":"wrong!",
			})
			return
		}
		id,err:=strconv.Atoi(Pid)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{
				"method":  "POST",
				"routing": "like",
				"Error":err.Error(),
			})
			return
		}else if id<0 {
			c.JSON(http.StatusOK, gin.H{
				"method":  "POST",
				"routing": "like",
				"Error":"illegal id !",
			})
			return
		}
		err=DCClient.Call("Message.Like",id,&DC.Message{})
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"like>Error:":err.Error(),
			})
			return
		}
		//Message.Like(id,db) //正式点赞
		//信息反馈
		var tmp []DC.Message
		err=DCClient.Call("Message.GetContent",DC.Message{},&tmp)
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"Get>Error:":err.Error(),
			})
			return
		}
		//tmp := Message.GetContent(db)
		c.JSON(http.StatusOK,&tmp)
	}
}
