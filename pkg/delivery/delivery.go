package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessNoContent(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

func Failed(c *gin.Context, httpCode int, err error) {
	resp := ErrorResponse{Message: err.Error()}
	c.JSON(httpCode, resp)
	c.Abort()
}
