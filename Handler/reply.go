package Handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"MessageBoard/Model"
)

func Reply(db *gorm.DB) gin.HandlerFunc {
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
			Model.Save(&reply,db)
			tmp:=Model.GetContent(db)
			c.JSON(http.StatusOK, &tmp)
		}
	}
}
