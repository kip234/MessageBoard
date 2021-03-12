package Handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"MessageBoard/Model"
)

func Publish(db *gorm.DB) gin.HandlerFunc {
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
			if tmp:=Model.Save(&speech,db);tmp!=nil{
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