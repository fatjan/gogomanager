package department

import (
	"github.com/fatjan/gogomanager/internal/dto"
)

type UseCase interface {
	PostDepartment(*dto.DepartmentRequest) (*dto.DepartmentResponse, error)
	UpdateDepartment(int, *dto.DepartmentRequest) (*dto.DepartmentResponse, error)
}
