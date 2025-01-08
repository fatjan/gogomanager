package dto

type User struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	UserImageUri    string `json:"usermageUri"`
	CompanyName     string `json:"companyName"`
	CompanyImageUri string `json:"companyImageUri"`
}

type UserRequest struct {
	UserID int `json:"id"`
}
