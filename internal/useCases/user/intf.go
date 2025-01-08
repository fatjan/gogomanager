package user

import "github.com/fatjan/gogomanager/internal/dto"

type UseCase interface {
	GetUser(*dto.UserRequest) (*dto.User, error)
}
