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
	Post(*models.Employee) (*models.Employee, error)
}
