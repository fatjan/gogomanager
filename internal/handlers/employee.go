package handlers

import (
	"net/http"
	"strconv"

	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/useCases/employee"
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
	limit := ginCtx.DefaultQuery("limit", "5")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 5
	}
	if limitInt > 100 {
		limitInt = 100
	}

	offset := ginCtx.DefaultQuery("offset", "0")
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		offsetInt = 0
	}
	if offsetInt < 0 {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Offset cannot be negative"})
		return
	}

	identityNumber := ginCtx.DefaultQuery("identityNumber", "")
	name := ginCtx.DefaultQuery("name", "")
	gender := dto.GenderType(ginCtx.DefaultQuery("gender", ""))
	departmentID := ginCtx.DefaultQuery("departmentId", "0")

	_, err = strconv.Atoi(departmentID)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid departmentId value"})
		return
	}

	employeeRequest := dto.EmployeeRequest{
		IdentityNumber: identityNumber,
		Name:           name,
		Gender:         gender,
		DepartmentID:   departmentID,
		Limit:          limitInt,
		Offset:         offsetInt,
	}

	employeeResponse, err := r.employeeUseCase.GetAllEmployee(&employeeRequest)
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

	err := r.employeeUseCase.DeleteByIdentityNumber(identityNumber)
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

	response, err := r.employeeUseCase.UpdateEmployee(identityNumber, &req)

	if err != nil {
		switch err.Error() {
		case "employee not found":
			ginCtx.JSON(http.StatusNotFound, gin.H{"error": "identity number not found"})
			return
		case "duplicate identity number":
			ginCtx.JSON(http.StatusConflict, gin.H{"error": "identity number already exists"})
			return
		case "identity number is required":
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "identity number is required"})
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
	employeeResponse, err := r.employeeUseCase.PostEmployee(&employeeRequest, managerId)
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
