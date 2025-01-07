package dto

import "github.com/google/uuid"

type Department struct {
	ID uuid.UUID `pg:"type:uuid,default:gen_random_uuid()"`
    Name string
}

type PostDepartmentRequest struct {
    Name string `validate:"required,min=4,max=33"`
}

type DepartmentResponse struct {
    DepartmentID int    `json:"departmentId"` 
    Name         string `json:"name"`
}

func GetDepartmentResponse(dept *Department) *DepartmentResponse {
    return &DepartmentResponse{
        DepartmentID: dept.ID, 
        Name:         dept.Name,
    }
}