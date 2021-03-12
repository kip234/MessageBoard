package Handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"MessageBoard/Model"
)

func Comment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context){
		var comment Model.Message
		err := c.ShouldBind(&comment)//绑定参数
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"method":  "POST",
				"routing": "comment",
				"Error":   err.Error(),
			})
			return
		}
		//回馈数据
		Model.Save(&comment,db)
		tmp := Model.GetContent(db)
		c.JSON(http.StatusOK,&tmp)
	}
}
