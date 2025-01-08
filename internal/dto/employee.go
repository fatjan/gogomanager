package dto

type GenderType string

const (
	male   GenderType = "male"
	female GenderType = "female"
)

type Employee struct {
	ID               int
	IdentityNumber   string     `json:"identityNumber"`
	Name             string     `json:"name"`
	Gender           GenderType `json:"gender"`
	DepartmentID     string     `json:"departmentId"`
	EmployeeImageURI string     `json:"employeeImageUri"`
}

type EmployeeRequest struct {
	IdentityNumber   string `validate:"min=5,max=33"`
	Name             string `validate:"min=4,max=33"`
	Gender           GenderType
	DepartmentID     string
	EmployeeImageURI string
	Limit            int
	Offset           int
}

type EmployeeResponse struct {
	IdentityNumber   string `json:"identityNumber"`
	Name             string `json:"name"`
	Gender           string `json:"gender"`
	DepartmentID     string `json:"departmentId"`
	EmployeeImageURI string `json:"employeeImageUri"`
}

type GetAllEmployeeResponse struct {
	Employees []*EmployeeResponse
}

type UpdateEmployeeRequest struct {
	IdentityNumber   string     `json:"identityNumber" validate:"min=5,max=33"`
	Name             string     `json:"name" validate:"min=5,max=33"`
	EmployeeImageURI string     `json:"employeeImageUri"`
	Gender           GenderType `json:"gender"`
	DepartmentID     string     `json:"departmentId" binding:"required"`
}

type UpdateEmployeeResponse struct {
	IdentityNumber   string `json:"identityNumber"`
	Name             string `json:"name"`
	EmployeeImageURI string `json:"employeeImageUri"`
	Gender           string `json:"gender"`
	DepartmentID     string `json:"departmentId"`
}
