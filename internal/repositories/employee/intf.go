package employee

import "github.com/fatjan/gogomanager/internal/dto"

type Repository interface {
	GetAll(*dto.EmployeeRequest) ([]*dto.Employee, error)
}
