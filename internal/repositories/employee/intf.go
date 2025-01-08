package employee

import (
	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/models"
)

type Repository interface {
	GetAll(*dto.EmployeeRequest) ([]*models.Employee, error)
	DeleteByIdentityNumber(identityNumber string) error
	UpdateEmployee(identityNumber string, request *models.UpdateEmployee) (*models.UpdateEmployee, error)
	FindByIdentityNumber(identityNumber string) (*models.IdentityNumberEmployee, error)
	CheckDuplicateIdentityNumber(currentIdentityNumber string) (string, error)
}
