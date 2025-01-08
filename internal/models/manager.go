package models

import "time"

type Manager struct {
	ID              int       `db:"id"`
	Email           string    `db:"email"`
	Password        string    `db:"password_hash"`
	Name            string    `db:"name"`
	UserImageUri    string    `db:"user_image_uri"`
	CompanyName     string    `db:"company_name"`
	CompanyImageUri string    `db:"company_image_uri"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}
