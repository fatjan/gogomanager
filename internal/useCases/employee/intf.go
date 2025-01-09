package employee

import (
	"github.com/fatjan/gogomanager/internal/dto"
)

type UseCase interface {
	GetAllEmployee(*dto.EmployeeRequest) (*dto.GetAllEmployeeResponse, error)
	DeleteByIdentityNumber(identityNumber string) error
	UpdateEmployee(identityNumber string, req *dto.UpdateEmployeeRequest) (*dto.UpdateEmployeeResponse, error)
}
