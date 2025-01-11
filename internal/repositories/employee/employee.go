package employee

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

	if filter.DepartmentID != "" {
		whereClause = append(whereClause, fmt.Sprintf("department_id = $%d", argCount))
		args = append(args, filter.DepartmentID)
		argCount++
	}

	whereStr := "WHERE " + strings.Join(whereClause, " AND ")

	query := fmt.Sprintf(`
					SELECT id, name, manager_id, identity_number, gender, department_id, employee_image_uri, created_at, updated_at
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

func (r *repository) Post(ctx context.Context, employee *models.Employee) (*models.Employee, error) {
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

	_, err := r.db.ExecContext(
		ctx,
		query,
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

func (r *repository) DeleteByIdentityNumber(ctx context.Context, identityNumber string, managerId int) error {
	result, err := r.db.ExecContext(ctx, "DELETE FROM employees WHERE identity_number = $1 AND manager_id = $2", identityNumber, managerId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("employee is not found")
	}

	return nil
}

func (r *repository) UpdateEmployee(ctx context.Context, identityNumber string, request *models.UpdateEmployee) (*models.UpdateEmployee, error) {
	var employee models.UpdateEmployee

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	paramCount := 1

	if request.IdentityNumber != "" {
		setValues = append(setValues, fmt.Sprintf("identity_number = $%d", paramCount))
		args = append(args, request.IdentityNumber)
		paramCount++
	}
	if request.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name = $%d", paramCount))
		args = append(args, request.Name)
		paramCount++
	}
	if request.EmployeeImageURI != "" {
		setValues = append(setValues, fmt.Sprintf("employee_image_uri = $%d", paramCount))
		args = append(args, request.EmployeeImageURI)
		paramCount++
	}
	if request.Gender != "" {
		setValues = append(setValues, fmt.Sprintf("gender = $%d", paramCount))
		args = append(args, request.Gender)
		paramCount++
	}
	if request.DepartmentID != 0 {
		setValues = append(setValues, fmt.Sprintf("department_id = $%d", paramCount))
		args = append(args, request.DepartmentID)
		paramCount++
	}

	setValues = append(setValues, "created_at = current_timestamp", "updated_at = current_timestamp")

	if len(setValues) == 2 {
		return nil, errors.New("no fields to update")
	}

	query := fmt.Sprintf(`
				UPDATE employees
				SET %s
				WHERE identity_number = $%d
				RETURNING 
						id,
						identity_number,
						name,
						employee_image_uri,
						gender,
						department_id`,
		strings.Join(setValues, ", "),
		paramCount,
	)

	args = append(args, identityNumber)

	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&employee.ID,
		&employee.IdentityNumber,
		&employee.Name,
		&employee.EmployeeImageURI,
		&employee.Gender,
		&employee.DepartmentID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("employee not found")
		}
		return nil, err
	}

	return &employee, nil
}

func (r *repository) FindByIdentityNumberWithManagerID(ctx context.Context, identityNumber string, managerId int) (*models.IdentityNumberEmployee, error) {
	employee := &models.IdentityNumberEmployee{}

	query := `SELECT identity_number, name, employee_image_uri, gender, department_id
        FROM employees 
        WHERE identity_number = $1`
	args := []interface{}{identityNumber}

	query += ` and manager_id = $2`
	args = append(args, managerId)

	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&employee.IdentityNumber,
		&employee.Name,
		&employee.EmployeeImageURI,
		&employee.Gender,
		&employee.DepartmentID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			employee.IdentityNumber = ""
			return employee, nil
		}
		return nil, err
	}

	return employee, nil
}

func (r *repository) FindByIdentityNumber(identityNumber string) (*models.IdentityNumberEmployee, error) {
	employee := &models.IdentityNumberEmployee{}

	err := r.db.QueryRow(`
        SELECT identity_number, name, employee_image_uri, gender, department_id
        FROM employees 
        WHERE identity_number = $1`,
		identityNumber,
	).Scan(
		&employee.IdentityNumber,
		&employee.Name,
		&employee.EmployeeImageURI,
		&employee.Gender,
		&employee.DepartmentID,
	)

	if err != nil {
		return nil, err
	}

	return employee, nil
}
