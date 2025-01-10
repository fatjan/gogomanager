package employee

import (
	"context"

	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/models"
)

type Repository interface {
	GetAll(*dto.EmployeeRequest) ([]*models.Employee, error)
	DeleteByIdentityNumber(ctx context.Context, identityNumber string) error
	UpdateEmployee(ctx context.Context, identityNumber string, request *models.UpdateEmployee) (*models.UpdateEmployee, error)
	FindByIdentityNumberWithDepartmentID(ctx context.Context, identityNumber string, department int) (*models.IdentityNumberEmployee, error)
	Post(*models.Employee) (*models.Employee, error)
	FindByIdentityNumber(identityNumber string) (*models.IdentityNumberEmployee, error)
}
