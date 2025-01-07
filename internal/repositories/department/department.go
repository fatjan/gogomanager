package department

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/fatjan/gogomanager/internal/models"
	"github.com/jmoiron/sqlx"
	"log"
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

func (r *repository) FindOneByID(id int) (*models.Department, error) {

	department := &models.Department{}

	err := r.db.QueryRow(
		"SELECT id, name FROM departments WHERE id = $1",
		id,
	).Scan(&department.ID, &department.Name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(fmt.Sprintf("department with id %d not found", id))
		}
		return nil, err
	}

	return department, nil

}

func (r *repository) Update(id int, department *models.Department) error {
	result, err := r.db.Exec(
		"UPDATE departments SET name = $1, manager_id = $2 WHERE id = $3",
		department.Name,
		1,
		id,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("error query")
		return err
	}
	if rowsAffected == 0 {
		log.Println("failed update department")
		return errors.New(fmt.Sprintf("department with id %d not found", id))
	}

	return nil
}
