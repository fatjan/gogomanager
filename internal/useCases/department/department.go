package department

import (
	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/repositories/department"
	"github.com/fatjan/gogomanager/internal/models"
)

type useCase struct {
	departmentRepository department.Repository
}

func NewUseCase(departmentRepository department.Repository) UseCase {
	return &useCase{departmentRepository: departmentRepository}
}

func (uc *useCase) PostDepartment(departmentRequest *dto.DepartmentRequest) (*dto.DepartmentResponse, error) {
	newDepartment := &models.Department{
		ID: "this is id",
		Name: departmentRequest.Name,
	}
	err := uc.departmentRepository.Post(newDepartment)
	if err != nil {
		return nil, err
	}

	departmentResponse := &dto.DepartmentResponse{
        DepartmentID: "this is id", 
        Name:         departmentRequest.Name,
    }

	return departmentResponse, nil
}