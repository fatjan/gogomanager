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
	Post(*models.Department) (int, error)
	Update(int, *models.Department) error
	FindOneByID(int) (*models.Department, error)
	DeleteByID(int) error
	FindAllWithFilter(
		ctx context.Context,
		filter DepartmentFilter,
		pagination pagination.Request,
	) ([]*models.Department, error)
}
