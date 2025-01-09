package employee

import (
	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/models"
)

type Repository interface {
	GetAll(*dto.EmployeeRequest) ([]*models.Employee, error)
	Post(*models.Employee) (*models.Employee, error)
}
