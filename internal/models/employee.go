package models

import "time"

type GenderType string

const (
	male   	GenderType = "male"
	female	GenderType = "female"
)

type Employee struct {
	ID 					int 
	IdentityNumber   	string  `db:"identity_number"`
    Name 				string 
	Gender 				string 
	DepartmentID 		string  `db:"department_id"`
	ManagerID			int 	`db:"manager_id"`
	EmployeeImageURI 	string  `db:"employee_image_uri"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}