package user

import (
	"context"
	"github.com/fatjan/gogomanager/internal/dto"
)

type UseCase interface {
	GetUser(*dto.UserRequest) (*dto.User, error)
	UpdateUser(context.Context, int, *dto.UserPatchRequest) (*dto.User, error)
}
