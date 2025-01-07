package handlers

import (
	"net/http"

	"github.com/fatjan/gogomanager/internal/useCases/department"
	"github.com/gin-gonic/gin"
)

type DepartmentHandler interface {
	Index(ginCtx *gin.Context)
	Detail(ginCtx *gin.Context)
}

type departmentHandler struct {
	departmentUseCase department.UseCase
}

func (r *departmentHandler) Index(ginCtx *gin.Context) {
	ginCtx.JSON(http.StatusOK, nil)
}

func (r *departmentHandler) Detail(ginCtx *gin.Context) {
	ginCtx.JSON(http.StatusOK, nil)
}

func NewDepartmentHandler(departmentUseCase department.UseCase) DepartmentHandler {
	return &departmentHandler{departmentUseCase: departmentUseCase}
}
