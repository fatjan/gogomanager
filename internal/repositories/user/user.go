package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/fatjan/gogomanager/internal/dto"
	"log"

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
	nullFields := struct {
		Name            sql.NullString
		UserImageUri    sql.NullString
		CompanyName     sql.NullString
		CompanyImageUri sql.NullString
	}{}

	user := &models.User{}

	err := r.db.QueryRow(`
	select 
		email, "name", user_image_uri, company_name, company_image_uri 
	from managers 
	where id = $1;`, id).Scan(
		&user.Email,
		&nullFields.Name,
		&nullFields.UserImageUri,
		&nullFields.CompanyName,
		&nullFields.CompanyImageUri,
	)

	if err != nil {
		return nil, err
	}

	user.Name = nullFields.Name.String
	user.UserImageUri = nullFields.UserImageUri.String
	user.CompanyName = nullFields.CompanyName.String
	user.CompanyImageUri = nullFields.CompanyImageUri.String

	return user, nil
}
func (r *repository) Update(_ context.Context, userID int, request *dto.UserPatchRequest) error {

	result, err := r.db.Exec(
		"UPDATE managers SET email = $1,name = $2,user_image_uri = $3, company_name = $4, company_image_uri = $5 WHERE id = $6",
		request.Email, request.Name, request.UserImageUri, request.CompanyName, request.CompanyImageUri,
		userID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("error query")
		return err
	}
	if rowsAffected == 0 {
		log.Println("failed update manager")
		return errors.New("update query failed")
	}

	return nil
}
