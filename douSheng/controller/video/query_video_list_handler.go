package video

import (
	"douSheng/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryVideoListHandler(c *gin.Context) {
	c.JSON(http.StatusOK,
		models.CommonResponse{
			StatusCode: 1,
			StatusMsg:  "火热开发中, 待3月8日期末考试后上线",
		})
}
