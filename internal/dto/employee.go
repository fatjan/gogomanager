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
	IdentityNumber   string     `validate:"required,min=5,max=33"`
	Name             string     `validate:"required,min=4,max=33"`
	Gender           GenderType `validate:"oneof=male f emale"`
	DepartmentID     string     `validate:"required"`
	EmployeeImageURI string     `validate:"string"`
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

type GetAllEmployeeRequest struct {
	IdentityNumber  string		`form:"identityNumber"`
	Name      		string 		`form:"name"`
	Gender      	string 		`form:"gender"`
	DepartmentID	string 		`form:"departmentId"`
	ManagerID 		int    		`form:"-"`

	PaginationRequest
}