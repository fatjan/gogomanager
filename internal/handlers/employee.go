package handlers

import (
	"net/http"

	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/useCases/employee"
	"github.com/fatjan/gogomanager/pkg/delivery"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler interface {
	Get(ginCtx *gin.Context)
	Post(ginCtx *gin.Context)
}

type employeeHandler struct {
	employeeUseCase employee.UseCase
}

func (r *employeeHandler) Get(ginCtx *gin.Context) {
	var req dto.GetAllEmployeeRequest

	if err := ginCtx.ShouldBindQuery(&req); err != nil {
		delivery.Failed(ginCtx, http.StatusBadRequest, err)
		return
	}

	managerId, exists := ginCtx.Get("manager_id")
	if !exists {
		ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid manager id"})
		return
	}
	id := managerId.(int)
	req.ManagerID = id

	employeeResponse, err := r.employeeUseCase.GetAllEmployee(ginCtx.Request.Context(), req)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, employeeResponse)
}

func (r *employeeHandler) Post(ginCtx *gin.Context) {
	var employeeRequest dto.EmployeeRequest
	if err := ginCtx.BindJSON(&employeeRequest); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	managerId := ginCtx.GetInt("manager_id")
	employeeResponse, err := r.employeeUseCase.PostEmployee(ginCtx.Request.Context(), &employeeRequest, managerId)
	if err != nil {
		if err.Error() == "duplicate identity number" {
			ginCtx.JSON(http.StatusConflict, gin.H{"error": "Duplicate identity number"})
			return
		}
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, employeeResponse)
}

func NewEmployeeHandler(employeeUseCase employee.UseCase) EmployeeHandler {
	return &employeeHandler{employeeUseCase: employeeUseCase}
}
