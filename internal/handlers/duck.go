package handlers

import (
	"net/http"

	"github.com/fatjan/gogomanager/internal/useCases/duck"
	"github.com/gin-gonic/gin"
)

type DuckHandler interface {
	Index(ginCtx *gin.Context)
	Detail(ginCtx *gin.Context)
}

type duckHandler struct {
	duckUseCase duck.UseCase
}

func (r *duckHandler) Index(ginCtx *gin.Context) {
	ginCtx.JSON(http.StatusOK, nil)
}

func (r *duckHandler) Detail(ginCtx *gin.Context) {
	ginCtx.JSON(http.StatusOK, nil)
}

func NewDuckHandler(duckUseCase duck.UseCase) DuckHandler {
	return &duckHandler{duckUseCase: duckUseCase}
}
