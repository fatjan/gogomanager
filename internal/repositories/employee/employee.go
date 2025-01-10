package employee

import (
	"errors"
	"fmt"
	"context"
	"strings"

	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const (
	PG_DUPLICATE_ERROR = "23505"
)

type repository struct {
	db *sqlx.DB
}

func NewEmployeeRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context, filter EmployeeFilter, page dto.PaginationRequest) ([]*models.Employee, error) {
	whereClause := []string{}
	args := []interface{}{}
	argCount := 1

	whereClause = append(whereClause, fmt.Sprintf("manager_id = $%d", argCount))
	args = append(args, filter.ManagerID)
	argCount++

	if filter.IdentityNumber != "" {
		whereClause = append(whereClause, fmt.Sprintf("identity_number ILIKE $%d", argCount))
		args = append(args, "%"+filter.IdentityNumber+"%")
		argCount++
	}

	if filter.Name != "" {
		whereClause = append(whereClause, fmt.Sprintf("name ILIKE $%d", argCount))
		args = append(args, "%"+filter.Name+"%")
		argCount++
	}

	if filter.Gender != "" {
		whereClause = append(whereClause, fmt.Sprintf("gender = $%d", argCount))
		args = append(args, filter.Gender)
		argCount++
	}

	if filter.DepartmentID != "0" {
		whereClause = append(whereClause, fmt.Sprintf("department_id = $%d", argCount))
		args = append(args, filter.Gender)
		argCount++
	}

	whereStr := "WHERE " + strings.Join(whereClause, " AND ")

	query := fmt.Sprintf(`
					SELECT id, name, manager_id, identity_number, gender, department_id, created_at, updated_at
					FROM employees
					%s
					ORDER BY id
					LIMIT $%d OFFSET $%d`,
		whereStr, argCount, argCount+1)

	args = append(args, page.GetLimit(), page.GetOffset())

	employees := []*models.Employee{}
	err := r.db.SelectContext(ctx, &employees, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error get all employees query: %w", err)
	}

	return employees, nil
}

func (r *repository) Post(employee *models.Employee) (*models.Employee, error) {
	query := `
			INSERT INTO employees (
				identity_number,
				name,
				employee_image_uri,
				gender,
				department_id,
				manager_id
			) VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(query,
		employee.IdentityNumber,
		employee.Name,
		employee.EmployeeImageURI,
		employee.Gender,
		employee.DepartmentID,
		employee.ManagerID,
	)

	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code == PG_DUPLICATE_ERROR {
			return nil, errors.New("duplicate identity number")
		}

		return nil, err
	}

	return employee, err
}
