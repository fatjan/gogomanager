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

func (r *repository) Post(department *models.Department) (int, error) {
	var newID int
	err := r.db.QueryRow("INSERT INTO departments (name, manager_id) VALUES ($1, $2) RETURNING id", department.Name, 1).Scan(&newID)
	if err != nil {
		return 0, err
	}
	
	return newID, nil
}
