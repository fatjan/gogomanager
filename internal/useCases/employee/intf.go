package employee

import (
	"context"

	"github.com/fatjan/gogomanager/internal/dto"
)

type UseCase interface {
	GetAllEmployee(context.Context, dto.GetAllEmployeeRequest) ([]*dto.EmployeeResponse, error)
	DeleteByIdentityNumber(ctx context.Context, identityNumber string) error
	UpdateEmployee(ctx context.Context, identityNumber string, req *dto.UpdateEmployeeRequest) (*dto.UpdateEmployeeResponse, error)
	PostEmployee(context.Context, *dto.EmployeeRequest, int) (*dto.EmployeeResponse, error)
}
