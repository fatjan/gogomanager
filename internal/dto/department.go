package dto

type Department struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DepartmentRequest struct {
	Name string `validate:"required,min=4,max=33"`
}

type DepartmentResponse struct {
	DepartmentID string `json:"departmentId"`
	Name         string `json:"name"`
}

type GetAllDepartmentRequest struct {
	Name      string `form:"name"`
	ManagerID int    `form:"-"`

	PaginationRequest
}
