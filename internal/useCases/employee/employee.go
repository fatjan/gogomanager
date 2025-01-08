package employee

import (
	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/repositories/employee"
)

type useCase struct {
	employeeRepository employee.Repository
}

func NewUseCase(employeeRepository employee.Repository) UseCase {
	return &useCase{employeeRepository: employeeRepository}
}

func (uc *useCase) GetAllEmployee(employeeRequest *dto.EmployeeRequest) (*dto.GetAllEmployeeResponse, error) {
	employees, err := uc.employeeRepository.GetAll(employeeRequest)
	if err != nil {
		return nil, err
	}

	allEmployee := make([]*dto.EmployeeResponse, 0)
	for _, v := range employees {
		employeeDto := &dto.EmployeeResponse{
			Name: v.Name,
			Gender: v.Gender,
			DepartmentID: v.DepartmentID,
			EmployeeImageURI: v. EmployeeImageURI,
			IdentityNumber: v.IdentityNumber,
		}
		allEmployee = append(allEmployee, employeeDto)
	}

	return &dto.GetAllEmployeeResponse{Employees: allEmployee}, nil
}