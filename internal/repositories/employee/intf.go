package employee

import (
	"context"

	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/models"
)

type EmployeeFilter struct {
	ManagerID 		int    	`db:"manager_id"`
	Name      		string 	`db:"name"`
	IdentityNumber  string	`db:"identity_number"`
	Gender      	string 	`db:"gender"`
	DepartmentID	string	`db:"department_id"`
}

type Repository interface {
	GetAll(context.Context,
		EmployeeFilter,
		dto.PaginationRequest,
	) ([]*models.Employee, error)
	DeleteByIdentityNumber(ctx context.Context, identityNumber string) error
	UpdateEmployee(ctx context.Context, identityNumber string, request *models.UpdateEmployee) (*models.UpdateEmployee, error)
	FindByIdentityNumberWithManagerID(ctx context.Context, identityNumber string, managerId int) (*models.IdentityNumberEmployee, error)
	Post(context.Context, *models.Employee) (*models.Employee, error)
	FindByIdentityNumber(identityNumber string) (*models.IdentityNumberEmployee, error)
}
