package department

import (
	"context"

	"github.com/fatjan/gogomanager/internal/dto"
)

type UseCase interface {
	PostDepartment(*dto.DepartmentRequest) (*dto.DepartmentResponse, error)
	UpdateDepartment(int, *dto.DepartmentRequest) (*dto.DepartmentResponse, error)
	DeleteDepartment(int) error
	GetAllDepartment(context.Context, dto.GetAllDepartmentRequest) (*dto.GetAllDepartmentResponse, error)
}
