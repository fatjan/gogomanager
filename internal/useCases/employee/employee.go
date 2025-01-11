package employee

import (
	"context"
	"errors"
	"strconv"

	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/models"
	"github.com/fatjan/gogomanager/internal/repositories/department"
	"github.com/fatjan/gogomanager/internal/repositories/employee"
)

type useCase struct {
	employeeRepository  employee.Repository
	department department.Repository
}

func NewUseCase(employeeRepository employee.Repository, departmentRepository department.Repository) UseCase {
	return &useCase{
		employeeRepository:  employeeRepository,
		department: departmentRepository,
	}
}

func (uc *useCase) GetAllEmployee(c context.Context, req dto.GetAllEmployeeRequest) ([]*dto.EmployeeResponse, error) {
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

func (uc *useCase) DeleteByIdentityNumber(c context.Context, identityNumber string, managerId int) error {
	err := uc.employeeRepository.DeleteByIdentityNumber(c, identityNumber, managerId)
	if err != nil {
		if err.Error() == "employee is not found" {
			return err
		}
	}

	return nil
}

func (uc *useCase) UpdateEmployee(c context.Context, identityNumber string, req *dto.UpdateEmployeeRequest) (*dto.UpdateEmployeeResponse, error) {
	var departmentID int = 0
	var err error

	if req.DepartmentID != "" {
		departmentID, err = strconv.Atoi(req.DepartmentID)
		if err != nil {
			return nil, errors.New("invalid department id format")
		}

		_, err = uc.department.FindOneByID(c, departmentID, req.ManagerID)
		if err != nil {
			if err.Error() == "department not found" {
				return nil, err
			}
		}
	}

	if req.IdentityNumber != "" {
		// Check if the identityNumber that the manager wants to update to already exists
		employeeExist, err := uc.employeeRepository.IdentityNumberExists(c, req.IdentityNumber, req.ManagerID)
		if err != nil {
			return nil, err
		}
		if employeeExist {
			return nil, errors.New("duplicate identity number")
		}
	}

	// ensure identityNumber being selected for patch exists
	exists, err := uc.employeeRepository.IdentityNumberExists(c, identityNumber, req.ManagerID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("employee not found") // If exists is false, handle the case of no employee found.
	}

	employees := models.UpdateEmployee{
		IdentityNumber:   req.IdentityNumber,
		Name:             req.Name,
		EmployeeImageURI: req.EmployeeImageURI,
		Gender:           models.GenderType(req.Gender),
		DepartmentID:     departmentID,
		ManagerID:     	  req.ManagerID,
	}

	updatedEmployee, err := uc.employeeRepository.UpdateEmployee(c, identityNumber, &employees)
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

func (uc *useCase) PostEmployee(c context.Context, employeeRequest *dto.EmployeeRequest, managerId int) (*dto.EmployeeResponse, error) {
	newEmployee := &models.Employee{
		Name:             employeeRequest.Name,
		IdentityNumber:   employeeRequest.IdentityNumber,
		Gender:           string(employeeRequest.Gender),
		DepartmentID:     employeeRequest.DepartmentID,
		EmployeeImageURI: employeeRequest.EmployeeImageURI,
		ManagerID:        managerId,
	}

	createdEmployee, err := uc.employeeRepository.Post(c, newEmployee)
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
