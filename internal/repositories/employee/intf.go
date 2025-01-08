package employee

import "github.com/fatjan/gogomanager/internal/dto"
import "github.com/fatjan/gogomanager/internal/models"

type Repository interface {
	GetAll(*dto.EmployeeRequest) ([]*models.Employee, error)
}
