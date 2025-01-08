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
		delivery.Failed(ginCtx, http.StatusBadRequest, errors.New("invalid input"))
		return
	}

	departmentResponse, err := r.departmentUseCase.PostDepartment(ginCtx.Request.Context(), &departmentRequest)
	if err != nil {
		delivery.Failed(ginCtx, http.StatusInternalServerError, err)
		return
	}

	delivery.Success(ginCtx, departmentResponse)
}

func (r *departmentHandler) Update(ginCtx *gin.Context) {
	var departmentRequest dto.DepartmentRequest

	departmentID := ginCtx.Param("id")

	departmentIDInt, err := strconv.Atoi(departmentID)
	if err != nil {
		delivery.Failed(ginCtx, http.StatusBadRequest, err)
		return
	}

	if err := ginCtx.BindJSON(&departmentRequest); err != nil {
		delivery.Failed(ginCtx, http.StatusBadRequest, errors.New("invalid input"))
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

	delivery.Success(ginCtx, departmentResponse)
}

func (r *departmentHandler) Delete(ginCtx *gin.Context) {
	departmentID := ginCtx.Param("id")

	departmentIDInt, err := strconv.Atoi(departmentID)
	if err != nil {
		delivery.Failed(ginCtx, http.StatusBadRequest, err)
		return
	}

	err = r.departmentUseCase.DeleteDepartment(ginCtx.Request.Context(), departmentIDInt)
	if err != nil {
		statusRes := http.StatusInternalServerError
		errorMessageRes := errors.New("internal server error")
		if errors.Is(err, sql.ErrNoRows) {
			log.Println(fmt.Sprintf("department with id %d not found", departmentIDInt))
			statusRes = http.StatusNotFound
			errorMessageRes = errors.New(fmt.Sprintf("department with id %d not found", departmentIDInt))
		}

		delivery.Failed(ginCtx, statusRes, errorMessageRes)
		return
	}

	delivery.SuccessNoContent(ginCtx)
}

func (r *departmentHandler) Index(ginCtx *gin.Context) {
	var req dto.GetAllDepartmentRequest
	if err := ginCtx.ShouldBindQuery(&req); err != nil {
		delivery.Failed(ginCtx, http.StatusBadRequest, err)
		return
	}

	// TODO: set literal value of managerID for now, but later should use value from authentication process
	req.ManagerID = 1

	response, err := r.departmentUseCase.GetAllDepartment(ginCtx.Request.Context(), req)
	if err != nil {
		delivery.Failed(ginCtx, http.StatusInternalServerError, err)
		return
	}

	delivery.Success(ginCtx, response)
}

func NewDepartmentHandler(departmentUseCase department.UseCase) DepartmentHandler {
	return &departmentHandler{departmentUseCase: departmentUseCase}
}
