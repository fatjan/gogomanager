package handlers

import (
	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/useCases/employee"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EmployeeHandler interface {
	Get(ginCtx *gin.Context)
}

type employeeHandler struct {
	employeeUseCase employee.UseCase
}

func (r *employeeHandler) Get(ginCtx *gin.Context) {
	limit := ginCtx.DefaultQuery("limit", "5")
	_, err := strconv.Atoi(limit)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	offset := ginCtx.DefaultQuery("offset", "0")
	_, err = strconv.Atoi(offset)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset value"})
		return
	}

	idNumber := ginCtx.DefaultQuery("identityNumber", "")
	name := ginCtx.DefaultQuery("name", "")
	gender := dto.GenderType(ginCtx.DefaultQuery("gender", ""))
	departmentID := ginCtx.DefaultQuery("departmentId", "0")
	employeeImageURI := ginCtx.DefaultQuery("employeeImageUri", "")

	_, err = strconv.Atoi(departmentID)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid departmentId value"})
		return
	}

	employeeRequest := dto.EmployeeRequest{
		IdentityNumber: idNumber,
		Name: name,
		Gender: gender,
		DepartmentID: departmentID,
		EmployeeImageURI: employeeImageURI,
		Limit: limit,
		Offset: offset,
	}

	employeeResponse, err := r.employeeUseCase.GetAllEmployee(&employeeRequest)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, employeeResponse)
}

func NewEmployeeHandler(employeeUseCase employee.UseCase) EmployeeHandler {
	return &employeeHandler{employeeUseCase: employeeUseCase}
}
