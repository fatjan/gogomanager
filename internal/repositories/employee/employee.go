package employee

import (
	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/models"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
)

type repository struct {
	db *sqlx.DB
}

func NewEmployeeRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(employeeRequest *dto.EmployeeRequest) ([]*models.Employee, error) {
	baseQuery := "SELECT * FROM employees WHERE 1=1"
	var args []interface{}
	var argIndex int

	idNumber := employeeRequest.IdentityNumber
	name := employeeRequest.Name
	gender := employeeRequest.Gender
	departmentID := employeeRequest.DepartmentID
	employeeImageURI := employeeRequest.EmployeeImageURI

	if idNumber != "" {
		baseQuery += " AND identity_number = $" + strconv.Itoa(argIndex+1)
		args = append(args, idNumber)
		argIndex++
	}

	if name != "" {
		baseQuery += " AND name = $" + strconv.Itoa(argIndex+1)
		args = append(args, name)
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

	if employeeImageURI != "" {
		baseQuery += " AND employee_image_uri = $" + strconv.Itoa(argIndex+1)
		args = append(args, employeeImageURI)
	}

	employees := make([]*models.Employee, 0)

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
