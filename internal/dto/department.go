package dto

import "github.com/google/uuid"

type Department struct {
	ID uuid.UUID `pg:"type:uuid,default:gen_random_uuid()"`
    Name string
}

type PostDepartmentRequest struct {
    Name string `validate:"required,min=4,max=33"`
}

type GetDepartmentResponse struct {
    DepartmentID int    `json:"departmentId"` 
    Name         string `json:"name"`
}

func NewGetDepartmentResponse(dept *Department) *GetDepartmentResponse {
    return &GetDepartmentResponse{
        DepartmentID: dept.ID, 
        Name:         dept.Name,
    }
}