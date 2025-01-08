package department

import (
	"context"

	"github.com/fatjan/gogomanager/internal/models"
	"github.com/fatjan/gogomanager/pkg/pagination"
)

type DepartmentFilter struct {
	ManagerID int    `db:"manager_id"`
	Name      string `db:"name"`
}

type Repository interface {
	Post(context.Context, *models.Department) (int, error)
	Update(context.Context, int, *models.Department) error
	FindOneByID(context.Context, int) (*models.Department, error)
	DeleteByID(context.Context, int) error
	FindAllWithFilter(
		context.Context,
		DepartmentFilter,
		pagination.Request,
	) ([]*models.Department, error)
}
