package employee

import (
	"context"

	"github.com/fatjan/gogomanager/internal/dto"
)

type UseCase interface {
	GetAllEmployee(context.Context, dto.GetAllEmployeeRequest) ([]*dto.EmployeeResponse, error)
	PostEmployee(context.Context, *dto.EmployeeRequest, int) (*dto.EmployeeResponse, error)
}
