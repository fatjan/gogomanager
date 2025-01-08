package auth

import (
	"log"
	"time"

	"github.com/fatjan/gogomanager/internal/models"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Post(manager *models.Manager) (int, error) {
	var newID int
	now := time.Now()
	err := r.db.QueryRow("INSERT INTO managers (email, password_hash, name, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id", manager.Email, manager.Password, manager.Name, now, now).Scan(&newID)
	if err != nil {
		log.Println("error query")
		return 0, err
	}

	return newID, nil
}

func (r *repository) FindByEmail(email string) (*models.Manager, error) {
	manager := &models.Manager{}

	err := r.db.QueryRow(
		"SELECT id, email, password_hash, name FROM managers WHERE email = $1",
		email,
	).Scan(&manager.ID, &manager.Email, &manager.Password, &manager.Name)

	if err != nil {
		return nil, err
	}

	return manager, nil

}
