package employee

import (
	"context"

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

func (uc *useCase) GetAllEmployee(c context.Context, req *dto.GetAllEmployeeRequest) ([]*dto.EmployeeResponse, error) {
	filter := employee.EmployeeFilter{
		ManagerID: req.ManagerID,
		Name: req.Name,
		Gender: req.Gender,
		IdentityNumber: req.IdentityNumber,
		DepartmentID: req.DepartmentID,
	}
	
	employees, err := uc.employeeRepository.GetAll(c, filter, req.PaginationRequest)
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

	return allEmployee, nil
}

func (uc *useCase) PostEmployee(c context.Context, employeeRequest *dto.EmployeeRequest, managerId int) (*dto.EmployeeResponse, error) {
	newEmployee := &models.Employee{
		Name:             employeeRequest.Name,
		IdentityNumber:   employeeRequest.IdentityNumber,
		Gender:           string(employeeRequest.Gender),
		DepartmentID:     employeeRequest.DepartmentID,
		EmployeeImageURI: employeeRequest.EmployeeImageURI,
		ManagerID:        managerId,
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
