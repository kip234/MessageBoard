package Handler

import (
	"MessageBoard/Model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Publish(db interface{}) gin.HandlerFunc {
	return func(c *gin.Context){
		var speech Model.Message
		//:=c.PostForm("Content")
		err := c.ShouldBind(&speech)
		if err != nil {
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
			if tmp:=speech.Save(db);tmp!=nil{
				c.JSON(http.StatusBadRequest, gin.H{
					"method":  "POST",
					"routing": "publish",
					"Error": tmp.Error(),
				})
				return
			}
		}

		tmp:=Model.GetContent(db)
		c.JSON(http.StatusOK,&tmp)
	}
}