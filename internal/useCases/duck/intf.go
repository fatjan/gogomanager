package duck

import (
	"github.com/fatjan/gogomanager/internal/dto"
)

type UseCase interface {
	GetAllDuck() (*dto.GetAllDuckResponse, error)
	GetDuck(id int) (*dto.Duck, error)
}
