package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/fatjan/gogomanager/internal/dto"
	"log"
	"strings"

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
func (r *repository) Update(ctx context.Context, userID int, request *dto.UserPatchRequest) error {

	baseQuery := `UPDATE managers SET `
	var setClauses []string
	var args []interface{}
	var argIndex int = 1

	if request != nil {
		if request.Email != nil {
			setClauses = append(setClauses, fmt.Sprintf(`email = $%d`, argIndex))
			args = append(args, *request.Email)
			argIndex++
		}

		if request.Name != nil {
			setClauses = append(setClauses, fmt.Sprintf(`name = $%d`, argIndex))
			args = append(args, *request.Name)
			argIndex++
		}

		if request.UserImageUri != nil {
			setClauses = append(setClauses, fmt.Sprintf(`user_image_uri = $%d`, argIndex))
			args = append(args, *request.UserImageUri)
			argIndex++
		}

		if request.CompanyName != nil {
			setClauses = append(setClauses, fmt.Sprintf(`company_name = $%d`, argIndex))
			args = append(args, *request.CompanyName)
			argIndex++
		}

		if request.CompanyImageUri != nil {
			setClauses = append(setClauses, fmt.Sprintf(`company_image_uri = $%d`, argIndex))
			args = append(args, *request.CompanyImageUri)
			argIndex++
		}
	}

	if len(setClauses) == 0 {
		return errors.New("no fields to update")
	}

	baseQuery += strings.Join(setClauses, ", ")
	baseQuery += fmt.Sprintf(` WHERE id = $%d`, argIndex)
	args = append(args, userID)

	result, err := r.db.ExecContext(ctx, baseQuery, args...)
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
