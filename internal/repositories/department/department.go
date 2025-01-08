package department

import (
	"errors"
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
		log.Println("error query")
		return 0, err
	}

	return newID, nil
}

func (r *repository) FindOneByID(id int) (*models.Department, error) {

	department := &models.Department{}

	err := r.db.QueryRow(
		"SELECT id, name, manager_id, created_at, updated_at FROM departments WHERE id = $1",
		id,
	).Scan(&department.ID, &department.Name, &department.ManagerID, &department.CreatedAt, &department.UpdatedAt)

	if err != nil {

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
		return errors.New("update query failed")
	}

	return nil
}

func (r *repository) DeleteByID(id int) error {
	result, err := r.db.Exec("DELETE FROM departments WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		log.Println("deleted query is failed")
		return errors.New("deleted query failed")
	}

	return nil
}
