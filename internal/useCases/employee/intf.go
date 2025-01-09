package employee

import (
	"github.com/fatjan/gogomanager/internal/dto"
)

type UseCase interface {
	GetAllEmployee(*dto.EmployeeRequest) (*dto.GetAllEmployeeResponse, error)
	DeleteByIdentityNumber(identityNumber string) error
	UpdateEmployee(identityNumber string, req *dto.UpdateEmployeeRequest) (*dto.UpdateEmployeeResponse, error)
	PostEmployee(*dto.EmployeeRequest, int) (*dto.EmployeeResponse, error)
}
