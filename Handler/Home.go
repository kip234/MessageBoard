package Handler

import (
	"MessageBoard/Model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(db interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		message:=Model.GetContent(db)
		//回馈数据
		c.JSON(http.StatusOK,&message)
	}
}
