package department

import (
<<<<<<< HEAD
	"errors"
=======
>>>>>>> refs/remotes/origin/main
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
<<<<<<< HEAD
=======
		log.Println("error query")
>>>>>>> refs/remotes/origin/main
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
<<<<<<< HEAD
		return errors.New(fmt.Sprintf("department with id %d not found", id))
=======
		return fmt.Errorf("department with id %d not found", id)
>>>>>>> refs/remotes/origin/main
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
		log.Println(fmt.Sprintf("department with id %d not found", id))
		return err
	}

	return nil
}
