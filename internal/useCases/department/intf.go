package department 

import (
	"github.com/fatjan/gogomanager/internal/dto/department"
)

type UseCase interface {
	GetDepartment() (*dto.GetDepartmentResponse, error)
	PostDepartment() (*dto.PostDepartmentRequest. error)
}