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
	"github.com/go-playground/validator/v10"
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
	if ginCtx.ContentType() != "application/json" {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content type"})
		return
	}

	var departmentRequest dto.DepartmentRequest
	if err := ginCtx.BindJSON(&departmentRequest); err != nil {
		delivery.Failed(ginCtx, http.StatusBadRequest, errors.New("invalid input"))
		return
	}

	var validate = validator.New()
	if err := validate.Struct(departmentRequest); err != nil {
		delivery.Failed(ginCtx, http.StatusBadRequest, errors.New("validation error"))
		return
	}

	managerId, exists := ginCtx.Get("manager_id")
	if !exists {
		ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid manager id"})
		return
	}
	id := managerId.(int)
	departmentRequest.ManagerID = id

	departmentResponse, err := r.departmentUseCase.PostDepartment(ginCtx.Request.Context(), &departmentRequest)
	if err != nil {
		delivery.Failed(ginCtx, http.StatusInternalServerError, err)
		return
	}

	delivery.Success(ginCtx, http.StatusCreated, departmentResponse)
}

func (r *departmentHandler) Update(ginCtx *gin.Context) {
	if ginCtx.ContentType() != "application/json" {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content type"})
		return
	}

	var departmentRequest dto.DepartmentRequest

	departmentID := ginCtx.Param("id")

	departmentIDInt, err := strconv.Atoi(departmentID)
	if err != nil {
		delivery.Failed(ginCtx, http.StatusNotFound, err)
		return
	}

	managerId, exists := ginCtx.Get("manager_id")
	if !exists {
		ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid manager id"})
		return
	}
	id := managerId.(int)
	departmentRequest.ManagerID = id

	if err := ginCtx.BindJSON(&departmentRequest); err != nil {
		delivery.Failed(ginCtx, http.StatusBadRequest, errors.New("invalid input"))
		return
	}

	var validate = validator.New()
	if err := validate.Struct(departmentRequest); err != nil {
		delivery.Failed(ginCtx, http.StatusBadRequest, errors.New("validation error"))
		return
	}

	departmentResponse, err := r.departmentUseCase.UpdateDepartment(ginCtx.Request.Context(), departmentIDInt, &departmentRequest)
	if err != nil {
		statusRes := http.StatusInternalServerError
		errorMessageRes := errors.New("internal server error")
		if errors.Is(err, sql.ErrNoRows) {
			statusRes = http.StatusNotFound
			errorMessageRes = errors.New(fmt.Sprintf("department with id %d not found", departmentIDInt))
		}

		delivery.Failed(ginCtx, statusRes, errorMessageRes)
		return
	}

	delivery.Success(ginCtx, http.StatusOK, departmentResponse)
}

func (r *departmentHandler) Delete(ginCtx *gin.Context) {
	departmentID := ginCtx.Param("id")

	departmentIDInt, err := strconv.Atoi(departmentID)
	if err != nil {
		delivery.Failed(ginCtx, http.StatusNotFound, err)
		return
	}

	managerID, exists := ginCtx.Get("manager_id")
	if !exists {
		ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid manager id"})
		return
	}
	managerIDInt := managerID.(int)

	err = r.departmentUseCase.DeleteDepartment(ginCtx.Request.Context(), departmentIDInt, managerIDInt)
	if err != nil {
		statusRes := http.StatusInternalServerError
		errorMessageRes := errors.New("internal server error")
		if errors.Is(err, sql.ErrNoRows) {
			log.Println(fmt.Sprintf("department with id %d not found", departmentIDInt))
			statusRes = http.StatusNotFound
			errorMessageRes = errors.New(fmt.Sprintf("department with id %d not found", departmentIDInt))
		}
		if err.Error() == "department still has employees" {
			statusRes = http.StatusConflict
			errorMessageRes = errors.New(fmt.Sprintf(err.Error()))
		}

		delivery.Failed(ginCtx, statusRes, errorMessageRes)
		return
	}

	delivery.SuccessNoContent(ginCtx, http.StatusOK)
}

func (r *departmentHandler) Index(ginCtx *gin.Context) {
	var req dto.GetAllDepartmentRequest
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

	response, err := r.departmentUseCase.GetAllDepartment(ginCtx.Request.Context(), req)
	if err != nil {
		delivery.Failed(ginCtx, http.StatusInternalServerError, err)
		return
	}

	delivery.Success(ginCtx, http.StatusOK, response)
}

func NewDepartmentHandler(departmentUseCase department.UseCase) DepartmentHandler {
	return &departmentHandler{departmentUseCase: departmentUseCase}
}
