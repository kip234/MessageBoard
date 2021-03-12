package Handler

import (
	"MessageBoard/Model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func Like(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context){
		Pid,ok:=c.GetPostForm("Pid")
		fmt.Println(Pid)
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
		Model.Like(uint(id),db)
		tmp :=Model.GetContent(db)
		c.JSON(http.StatusOK,&tmp)
	}
}
