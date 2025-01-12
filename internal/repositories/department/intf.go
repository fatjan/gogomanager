package department

import (
	"context"

	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/models"
)

type DepartmentFilter struct {
	ManagerID int    `db:"manager_id"`
	Name      string `db:"name"`
}

type Repository interface {
	Post(context.Context, *models.Department) (int, error)
	Update(context.Context, int, *models.Department) error
	FindOneByID(context.Context, int, int) (*models.Department, error)
	DeleteByID(context.Context, int, int) error
	FindAllWithFilter(
		context.Context,
		DepartmentFilter,
		dto.PaginationRequest,
	) ([]*models.Department, error)
	DepartmentHasEmployee(ctx context.Context, departmentId int, managerId int) (bool, error)
}
