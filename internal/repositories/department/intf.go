package department

import "github.com/fatjan/gogomanager/internal/models"

type Repository interface {
	Post() (*models.Duck, error)
}
