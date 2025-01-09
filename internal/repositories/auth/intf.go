package auth

import (
	"github.com/fatjan/gogomanager/internal/models"
)

type Repository interface {
	FindByEmail(email string) (*models.Manager, error)
	Post(payload *models.Manager) (int, error)
}
