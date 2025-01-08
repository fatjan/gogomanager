package models

type GenderType string

const (
	male   	GenderType = "male"
	female	GenderType = "female"
)

type Employee struct {
	ID 				int 
	IDNumber   		string `json:"id"`
    Name 			string `json:"name"`
	Gender 			GenderType `json:"gender"`
	DepartmentID 	string `json:"departmentId"`
	EmployeeImageURI 	string `json:"employeeImageUri"`
}