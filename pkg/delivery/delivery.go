package delivery

import (
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, httpCode int, data any) {
	c.JSON(httpCode, data)
}

func Failed(c *gin.Context, httpCode int, err error) {
	resp := ErrorResponse{Message: err.Error()}
	c.JSON(httpCode, resp)
	c.Abort()
}
