package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/useCases/department"
	"github.com/fatjan/gogomanager/pkg/delivery"
	"github.com/gin-gonic/gin"
)

type DepartmentHandler interface {
	Post(ginCtx *gin.Context)
	Update(ginCtx *gin.Context)
	Delete(ginCtx *gin.Context)
	Index(ginCtx *gin.Context)
}

type departmentHandler struct {
	departmentUseCase department.UseCase
}

func (r *departmentHandler) Post(ginCtx *gin.Context) {
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

func (r *departmentHandler) Update(ginCtx *gin.Context) {
	var departmentRequest dto.DepartmentRequest

	departmentID := ginCtx.Param("id")

	departmentIDInt, _ := strconv.Atoi(departmentID)

	if err := ginCtx.BindJSON(&departmentRequest); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	departmentResponse, err := r.departmentUseCase.UpdateDepartment(departmentIDInt, &departmentRequest)
	if err != nil {

		statusRes := http.StatusInternalServerError
		errorMessageRes := errors.New("internal server error")
		if errors.Is(err, sql.ErrNoRows) {
			log.Println(fmt.Sprintf("department with id %d not found", departmentIDInt))
			statusRes = http.StatusNotFound
			errorMessageRes = errors.New(fmt.Sprintf("department with id %d not found", departmentIDInt))
		}

		ginCtx.JSON(statusRes, gin.H{"error": errorMessageRes.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, departmentResponse)
}

func (r *departmentHandler) Delete(ginCtx *gin.Context) {
	departmentID := ginCtx.Param("id")

	departmentIDInt, _ := strconv.Atoi(departmentID)
	err := r.departmentUseCase.DeleteDepartment(departmentIDInt)
	if err != nil {

		statusRes := http.StatusInternalServerError
		errorMessageRes := errors.New("internal server error")
		if errors.Is(err, sql.ErrNoRows) {
			log.Println(fmt.Sprintf("department with id %d not found", departmentIDInt))
			statusRes = http.StatusNotFound
			errorMessageRes = errors.New(fmt.Sprintf("department with id %d not found", departmentIDInt))

		}

		ginCtx.JSON(statusRes, gin.H{"error": errorMessageRes.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, "OK")
}

func (r *departmentHandler) Index(ginCtx *gin.Context) {
	var req dto.GetAllDepartmentRequest
	if err := ginCtx.ShouldBindUri(&req); err != nil {
		delivery.Failed(ginCtx, http.StatusBadRequest, err)
		return
	}

	response, err := r.departmentUseCase.GetAllDepartment(ginCtx.Request.Context(), req)
	if err != nil {
		delivery.Failed(ginCtx, http.StatusInternalServerError, err)
		return
	}

	delivery.SuccessWithMetadata(ginCtx, response.Departments, response.PaginationResponse)
}

func NewDepartmentHandler(departmentUseCase department.UseCase) DepartmentHandler {
	return &departmentHandler{departmentUseCase: departmentUseCase}
}
