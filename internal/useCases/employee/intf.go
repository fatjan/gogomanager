package employee

import (
	"context"

	"github.com/fatjan/gogomanager/internal/dto"
)

type UseCase interface {
	GetAllEmployee(*dto.EmployeeRequest) (*dto.GetAllEmployeeResponse, error)
	DeleteByIdentityNumber(ctx context.Context, identityNumber string) error
	UpdateEmployee(ctx context.Context, identityNumber string, req *dto.UpdateEmployeeRequest) (*dto.UpdateEmployeeResponse, error)
	PostEmployee(*dto.EmployeeRequest, int) (*dto.EmployeeResponse, error)
}
