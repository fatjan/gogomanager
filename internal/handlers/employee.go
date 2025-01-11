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
	Delete(ginCtx *gin.Context)
	Update(ginCtx *gin.Context)
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

func (r *employeeHandler) Delete(ginCtx *gin.Context) {
	identityNumber := ginCtx.Param("identityNumber")
	if identityNumber == "" {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "identity number is required"})
		return
	}

	managerIdStr, exists := ginCtx.Get("manager_id")
	if !exists {
		ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid manager id"})
		return
	}
	managerId := managerIdStr.(int)

	err := r.employeeUseCase.DeleteByIdentityNumber(ginCtx.Request.Context(), identityNumber, managerId)
	if err != nil {
		if err.Error() == "employee is not found" {
			ginCtx.JSON(http.StatusNotFound, gin.H{"error": "identityNumber is not found"})
			return
		}
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{"message": "employee deleted successfully"})
}

func (r *employeeHandler) Update(ginCtx *gin.Context) {
	identityNumber := ginCtx.Param("identityNumber")
	if identityNumber == "" {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "identity number is required"})
		return
	}

	var req dto.UpdateEmployeeRequest
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "validation error"})
		return
	}

	managerId, exists := ginCtx.Get("manager_id")
	if !exists {
		ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid manager id"})
		return
	}
	id := managerId.(int)
	req.ManagerID = id

	response, err := r.employeeUseCase.UpdateEmployee(ginCtx.Request.Context(), identityNumber, &req)

	if err != nil {
		switch err.Error() {
		case "employee not found":
			ginCtx.JSON(http.StatusNotFound, gin.H{"error": "identity number not found"})
			return
		case "duplicate identity number":
			ginCtx.JSON(http.StatusConflict, gin.H{"error": "identity number already exists"})
			return
		case "department not found":
			ginCtx.JSON(http.StatusNotFound, gin.H{"error": "department not found"})
			return
		default:
			ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
	}
	ginCtx.JSON(http.StatusOK, response)
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
