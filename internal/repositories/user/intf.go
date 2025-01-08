package user

import "github.com/fatjan/gogomanager/internal/models"

type Repository interface {
	GetUser(id int) (*models.User, error)
}
