package department

import (
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

func (uc *useCase) PostDepartment(departmentRequest *dto.DepartmentRequest) (*dto.DepartmentResponse, error) {
	newDepartment := &models.Department{
		Name: departmentRequest.Name,
	}
	departmentId, err := uc.departmentRepository.Post(newDepartment)
	if err != nil {
		return nil, err
	}

	departmentResponse := &dto.DepartmentResponse{
		DepartmentID: departmentId,
		Name:         departmentRequest.Name,
	}

	return departmentResponse, nil
}

func (uc *useCase) UpdateDepartment(departmentID int, departmentRequest *dto.DepartmentRequest) (*dto.DepartmentResponse, error) {
	newDepartment := &models.Department{
		Name: departmentRequest.Name,
	}

	_, err := uc.departmentRepository.FindOneByID(departmentID)
	if err != nil {
		return nil, err
	}

	err = uc.departmentRepository.Update(departmentID, newDepartment)
	if err != nil {
		return nil, err
	}

	departmentResponse := &dto.DepartmentResponse{
		DepartmentID: departmentID,
		Name:         departmentRequest.Name,
	}

	return departmentResponse, nil
}
