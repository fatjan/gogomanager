package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessNoContent(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}

func SuccessWithMetadata(c *gin.Context, data any, metadata any) {
	r := Response{Data: data}

	if metadata != nil {
		r.Metadata = metadata
	}

	c.JSON(http.StatusOK, r)
}

func Failed(c *gin.Context, httpCode int, err error) {
	resp := ErrorResponse{Message: err.Error()}
	c.JSON(httpCode, resp)
	c.Abort()
}
