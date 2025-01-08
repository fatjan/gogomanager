package department

import "github.com/fatjan/gogomanager/internal/models"

type Repository interface {
	Post(*models.Department) (int, error)
	Update(int, *models.Department) error
	FindOneByID(int) (*models.Department, error)
	DeleteByID(int) error
}
