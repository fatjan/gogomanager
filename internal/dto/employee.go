package dto

import (
	"github.com/fatjan/gogomanager/internal/models"
)

type GenderType string

const (
	male   	GenderType = "male"
	female	GenderType = "female"
)

type Employee struct {
	ID 				int 
	IDNumber   		string `json:"identityNumber"`
    Name 			string `json:"name"`
	Gender 			GenderType `json:"gender"`
	DepartmentID 	string `json:"departmentId"`
	EmployeeImageURI 	string `json:"employeeImageUri"`
}

type EmployeeRequest struct {
	IdentityNumber   	string `validate:"min=5,max=33"`
    Name 				string `validate:"min=4,max=33"`
	Gender 				GenderType
	DepartmentID 		string
	EmployeeImageURI 	string 
	Limit 				string 
	Offset 				string
}

type EmployeeResponse struct {
	IDNumber   			string `json:"departmentId"`
    Name 				string `json:"name"`
	Gender 				models.GenderType `json:"gender"`
	DepartmentID 		string `json:"departmentId"`
	EmployeeImageURI 	string `json:"employeeImageUri"`
}

type GetAllEmployeeResponse struct {
	Employees []*EmployeeResponse
}
