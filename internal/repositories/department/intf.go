package department

import "github.com/fatjan/gogomanager/internal/models"

type Repository interface {
	Post(*models.Department) (int, error)
	Update(id int, department *models.Department) error
	FindOneByID(id int) (*models.Department, error)
}
