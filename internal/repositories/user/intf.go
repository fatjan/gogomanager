package user

import (
	"context"
	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/models"
)

type Repository interface {
	GetUser(id int) (*models.User, error)
	Update(context.Context, int, *dto.UserPatchRequest) error
}
