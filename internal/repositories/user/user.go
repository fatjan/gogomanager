package user

import (
	"github.com/fatjan/gogomanager/internal/models"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetUser(id int) (*models.User, error) {
	user := &models.User{}

	err := r.db.QueryRow(`
	select 
		email, "name", user_image_uri, company_name, company_image_uri 
	from managers 
	where id = $1;`, id).Scan(
		&user.Email,
		&user.Name,
		&user.UserImageUri,
		&user.CompanyName,
		&user.CompanyImageUri,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
