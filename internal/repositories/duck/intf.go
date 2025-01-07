package duck

import "github.com/fatjan/gogomanager/internal/models"

type Repository interface {
	GetAll() ([]*models.Duck, error)
	GetByID(id int) (*models.Duck, error)
}
