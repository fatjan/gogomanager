package department

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/fatjan/gogomanager/internal/models"
	"github.com/fatjan/gogomanager/pkg/pagination"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewDepartmentRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Post(ctx context.Context, department *models.Department) (int, error) {
	var newID int
	err := r.db.QueryRowContext(ctx,
		"INSERT INTO departments (name, manager_id) VALUES ($1, $2) RETURNING id",
		department.Name, 1).Scan(&newID)
	if err != nil {
		log.Println("error query")
		return 0, err
	}

	return newID, nil
}

func (r *repository) FindOneByID(ctx context.Context, id int) (*models.Department, error) {
	department := &models.Department{}

	err := r.db.QueryRowContext(ctx,
		"SELECT id, name, manager_id, created_at, updated_at FROM departments WHERE id = $1",
		id,
	).Scan(&department.ID, &department.Name, &department.ManagerID, &department.CreatedAt, &department.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return department, nil
}

func (r *repository) Update(ctx context.Context, id int, department *models.Department) error {
	result, err := r.db.ExecContext(ctx,
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

func (r *repository) DeleteByID(ctx context.Context, id int) error {
	result, err := r.db.ExecContext(ctx, "DELETE FROM departments WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("deleted query failed")
	}

	return nil
}

func (r *repository) FindAllWithFilter(ctx context.Context, filter DepartmentFilter, page pagination.Request) ([]*models.Department, error) {
	whereClause := []string{}
	args := []interface{}{}
	argCount := 1

	whereClause = append(whereClause, fmt.Sprintf("manager_id = $%d", argCount))
	args = append(args, filter.ManagerID)
	argCount++

	if filter.Name != "" {
		whereClause = append(whereClause, fmt.Sprintf("name ILIKE $%d", argCount))
		args = append(args, "%"+filter.Name+"%")
		argCount++
	}

	whereStr := "WHERE " + strings.Join(whereClause, " AND ")

	query := fmt.Sprintf(`
					SELECT id, name, manager_id, created_at, updated_at
					FROM departments
					%s
					ORDER BY id
					LIMIT $%d OFFSET $%d`,
		whereStr, argCount, argCount+1)

	args = append(args, page.GetLimit(), page.GetOffset())

	departments := []*models.Department{}
	err := r.db.SelectContext(ctx, &departments, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error get all department query: %w", err)
	}

	return departments, nil
}
