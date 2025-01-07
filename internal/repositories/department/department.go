package department

import (
	"github.com/fatjan/gogomanager/internal/models"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewDepartmentRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Post(department *models.Department) (error) {
	_, err := r.db.Exec("INSERT INTO departments (name) VALUES ($1)", department.Name)
	if err != nil {
		return err
	}
	
	return nil
}
