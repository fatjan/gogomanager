package handlers

import (
	"net/http"
	"github.com/fatjan/gogomanager/internal/useCases/department"
	"github.com/fatjan/gogomanager/internal/dto"
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
	var departmentRequest dto.DepartmentRequest
	if err := ginCtx.BindJSON(&departmentRequest); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	
	departmentResponse, err := r.departmentUseCase.PostDepartment(&departmentRequest)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, departmentResponse)
}

func (r *departmentHandler) Detail(ginCtx *gin.Context) {
	ginCtx.JSON(http.StatusOK, nil)
}

func NewDepartmentHandler(departmentUseCase department.UseCase) DepartmentHandler {
	return &departmentHandler{departmentUseCase: departmentUseCase}
}
