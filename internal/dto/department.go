package dto

type Department struct {
	ID   string `json:"id"`
    Name string `json:"name"`
}

type DepartmentRequest struct {
    Name string `validate:"required,min=4,max=33"`
}

type DepartmentResponse struct {
    DepartmentID string `json:"departmentId"` 
    Name         string `json:"name"`
}