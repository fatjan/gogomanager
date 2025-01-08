package employee

import (
	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/models"
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
			Name:             v.Name,
			Gender:           v.Gender,
			DepartmentID:     v.DepartmentID,
			EmployeeImageURI: v.EmployeeImageURI,
			IdentityNumber:   v.IdentityNumber,
		}
		allEmployee = append(allEmployee, employeeDto)
	}

	return &dto.GetAllEmployeeResponse{Employees: allEmployee}, nil
}

func (uc *useCase) PostEmployee(employeeRequest *dto.EmployeeRequest) (*dto.EmployeeResponse, error) {
	newEmployee := &models.Employee{
		Name:             employeeRequest.Name,
		IdentityNumber:   employeeRequest.IdentityNumber,
		Gender:           string(employeeRequest.Gender),
		DepartmentID:     employeeRequest.DepartmentID,
		EmployeeImageURI: employeeRequest.EmployeeImageURI,
		ManagerID:        1, // Hardcoded for now since we don't have the auth feature yet
	}

	createdEmployee, err := uc.employeeRepository.Post(newEmployee)
	if err != nil {
		return nil, err
	}

	employeeResponse := &dto.EmployeeResponse{
		Name:             createdEmployee.Name,
		IdentityNumber:   createdEmployee.IdentityNumber,
		Gender:           createdEmployee.Gender,
		DepartmentID:     createdEmployee.DepartmentID,
		EmployeeImageURI: createdEmployee.EmployeeImageURI,
	}

	return employeeResponse, nil
}
