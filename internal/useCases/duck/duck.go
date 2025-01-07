package duck

import (
	"fmt"

	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/repositories/duck"
)

type useCase struct {
	duckRepository duck.Repository
}

func NewUseCase(duckRepository duck.Repository) UseCase {
	return &useCase{duckRepository: duckRepository}
}

func (uc *useCase) GetAllDuck() (*dto.GetAllDuckResponse, error) {
	ducks, err := uc.duckRepository.GetAll()
	if err != nil {
		return nil, err
	}

	allDuck := make([]*dto.Duck, 0)
	for _, v := range ducks {
		duckDto := &dto.Duck{
			ID:   fmt.Sprint(v.ID),
			Name: v.Name,
		}
		allDuck = append(allDuck, duckDto)
	}

	return &dto.GetAllDuckResponse{Ducks: allDuck}, nil
}

func (uc *useCase) GetDuck(id int) (*dto.Duck, error) {
	duck, err := uc.duckRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.Duck{ID: fmt.Sprint(duck.ID), Name: duck.Name}, nil
}
