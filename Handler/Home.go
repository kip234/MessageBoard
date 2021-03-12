package Handler

import (
	"MessageBoard/Model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Home(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		message:=Model.GetContent(db)
		//回馈数据
		c.JSON(http.StatusOK,&message)
	}
}
