package dto

type User struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	UserImageUri    string `json:"userImageUri"`
	CompanyName     string `json:"companyName"`
	CompanyImageUri string `json:"companyImageUri"`
}

type UserRequest struct {
	UserID int `json:"id"`
}

type UserPatchRequest struct {
	Email           string `json:"email" validate:"email"`
	Name            string `json:"name" validate:"min=4,max=52"`
	UserImageUri    string `json:"userImageUri" validate:"uri"`
	CompanyName     string `json:"companyName" validate:"min=4,max=52"`
	CompanyImageUri string `json:"companyImageUri" validate:"uri"`
}
