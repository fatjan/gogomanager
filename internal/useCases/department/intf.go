package department

import (
	"context"

	"github.com/fatjan/gogomanager/internal/dto"
)

type UseCase interface {
	PostDepartment(context.Context, *dto.DepartmentRequest) (*dto.DepartmentResponse, error)
	UpdateDepartment(context.Context, int, *dto.DepartmentRequest) (*dto.DepartmentResponse, error)
	DeleteDepartment(context.Context, int) error
	GetAllDepartment(context.Context, dto.GetAllDepartmentRequest) ([]*dto.DepartmentResponse, error)
}
