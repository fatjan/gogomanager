package employee

import (
	"errors"
	"strconv"

	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/models"
	"github.com/fatjan/gogomanager/internal/repositories/department"
	"github.com/fatjan/gogomanager/internal/repositories/employee"
)

type useCase struct {
	employeeRepository  employee.Repository
	deparmentRepository department.Repository
}

func NewUseCase(employeeRepository employee.Repository, departmentRepository department.Repository) UseCase {
	return &useCase{
		employeeRepository:  employeeRepository,
		deparmentRepository: departmentRepository,
	}
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

func (uc *useCase) DeleteByIdentityNumber(identityNumber string) error {
	err := uc.employeeRepository.DeleteByIdentityNumber(identityNumber)
	if err != nil {
		if err.Error() == "employee is not found" {
			return err
		}
	}

	return nil
}

func (uc *useCase) UpdateEmployee(identityNumber string, req *dto.UpdateEmployeeRequest) (*dto.UpdateEmployeeResponse, error) {
	var departmentID int = 0

	if req.DepartmentID != "" {
		departmentID, err := strconv.Atoi(req.DepartmentID)
		if err != nil {
			return nil, errors.New("invalid department id format")
		}

		_, err = uc.deparmentRepository.FindOneByID(departmentID)
		if err != nil {
			if err.Error() == "department not found" {
				return nil, err
			}
		}
	}

	if req.IdentityNumber != "" {
		identityNumber = req.IdentityNumber
	}

	employee, err := uc.employeeRepository.FindByIdentityNumberWithDepartmentID(identityNumber, departmentID)
	if err != nil {
		return nil, err
	}

	if employee.IdentityNumber == req.IdentityNumber {
		return nil, errors.New("duplicate identity number")
	}

	employees := models.UpdateEmployee{
		IdentityNumber:   req.IdentityNumber,
		Name:             req.Name,
		EmployeeImageURI: req.EmployeeImageURI,
		Gender:           models.GenderType(req.Gender),
		DepartmentID:     departmentID,
	}

	updatedEmployee, err := uc.employeeRepository.UpdateEmployee(identityNumber, &employees)
	if err != nil {
		return nil, err
	}

	responseDepartmentID := strconv.Itoa(updatedEmployee.DepartmentID)

	response := &dto.UpdateEmployeeResponse{
		IdentityNumber:   updatedEmployee.IdentityNumber,
		Name:             updatedEmployee.Name,
		EmployeeImageURI: updatedEmployee.EmployeeImageURI,
		Gender:           string(updatedEmployee.Gender),
		DepartmentID:     responseDepartmentID,
	}

	return response, nil
}

func (uc *useCase) PostEmployee(employeeRequest *dto.EmployeeRequest, managerId int) (*dto.EmployeeResponse, error) {
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
