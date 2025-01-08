package employee

import (
	"errors"
	"fmt"
	"log"
	"strconv"

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

func (r *repository) GetAll(employeeRequest *dto.EmployeeRequest) ([]*models.Employee, error) {
	limit := employeeRequest.Limit
	offset := employeeRequest.Offset

	baseQuery := fmt.Sprintf("SELECT * FROM employees WHERE 1=1")
	var args []interface{}
	var argIndex int

	idNumber := employeeRequest.IdentityNumber
	name := employeeRequest.Name
	gender := employeeRequest.Gender
	departmentID := employeeRequest.DepartmentID

	if idNumber != "" {
		baseQuery += " AND identity_number ILIKE $" + strconv.Itoa(argIndex+1)
		args = append(args, "%"+idNumber+"%")
		argIndex++
	}

	if name != "" {
		baseQuery += " AND name ILIKE $" + strconv.Itoa(argIndex+1)
		args = append(args, "%"+name+"%")
		argIndex++
	}

	if gender != "" {
		baseQuery += " AND gender = $" + strconv.Itoa(argIndex+1)
		args = append(args, gender)
		argIndex++
	}

	if departmentID != "0" {
		baseQuery += " AND department_id = $" + strconv.Itoa(argIndex+1)
		args = append(args, departmentID)
		argIndex++
	}

	baseQuery += " LIMIT $" + strconv.Itoa(argIndex+1) + " OFFSET $" + strconv.Itoa(argIndex+2)
	args = append(args, limit, offset)

	employees := make([]*models.Employee, 0)

	log.Println(baseQuery)
	rows, err := r.db.Queryx(baseQuery, args...)
	if err != nil {
		log.Println("error query GetAll Employee")
		return nil, err
	}
	for rows.Next() {
		var employee models.Employee
		err := rows.StructScan(&employee)
		if err != nil {
			log.Println("error query select")
			return nil, err
		}
		employees = append(employees, &employee)
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
