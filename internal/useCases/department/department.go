package department

import (
	"context"
	"fmt"

	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/models"
	"github.com/fatjan/gogomanager/internal/repositories/department"
)

type useCase struct {
	departmentRepository department.Repository
}

func NewUseCase(departmentRepository department.Repository) UseCase {
	return &useCase{departmentRepository: departmentRepository}
}

func (uc *useCase) PostDepartment(c context.Context, departmentRequest *dto.DepartmentRequest) (*dto.DepartmentResponse, error) {
	newDepartment := &models.Department{
		Name:      departmentRequest.Name,
		ManagerID: departmentRequest.ManagerID,
	}
	departmentId, err := uc.departmentRepository.Post(c, newDepartment)
	if err != nil {
		return nil, err
	}

	departmentResponse := &dto.DepartmentResponse{
		DepartmentID: fmt.Sprint(departmentId),
		Name:         departmentRequest.Name,
	}

	return departmentResponse, nil
}

func (uc *useCase) UpdateDepartment(c context.Context, departmentID int, departmentRequest *dto.DepartmentRequest) (*dto.DepartmentResponse, error) {
	newDepartment := &models.Department{
		Name: departmentRequest.Name,
		ManagerID: departmentRequest.ManagerID,
	}

	_, err := uc.departmentRepository.FindOneByID(c, departmentID, departmentRequest.ManagerID)
	if err != nil {
		return nil, err
	}

	err = uc.departmentRepository.Update(c, departmentID, newDepartment)
	if err != nil {
		return nil, err
	}

	departmentResponse := &dto.DepartmentResponse{
		DepartmentID: fmt.Sprint(departmentID),
		Name:         departmentRequest.Name,
	}

	return departmentResponse, nil
}

func (uc *useCase) DeleteDepartment(c context.Context, departmentID int, managerID int) error {
	_, err := uc.departmentRepository.FindOneByID(c, departmentID, managerID)
	if err != nil {
		return err
	}

	err = uc.departmentRepository.DeleteByID(c, departmentID)
	if err != nil {
		return err
	}

	return nil
}

func (uc *useCase) GetAllDepartment(c context.Context, req dto.GetAllDepartmentRequest) ([]*dto.DepartmentResponse, error) {
	filter := department.DepartmentFilter{
		ManagerID: req.ManagerID,
		Name:      req.Name,
	}

	departmentRecords, err := uc.departmentRepository.FindAllWithFilter(c, filter, req.PaginationRequest)
	if err != nil {
		return nil, err
	}

	departments := make([]*dto.DepartmentResponse, 0)
	for _, v := range departmentRecords {
		departments = append(departments, &dto.DepartmentResponse{
			DepartmentID: fmt.Sprint(v.ID),
			Name:         v.Name,
		})
	}

	return departments, nil
}
