package employee

import (
	"github.com/fatjan/gogomanager/internal/dto"
)

type UseCase interface {
	GetAllEmployee(*dto.EmployeeRequest) (*dto.GetAllEmployeeResponse, error)
	PostEmployee(*dto.EmployeeRequest) (*dto.EmployeeResponse, error)
}
