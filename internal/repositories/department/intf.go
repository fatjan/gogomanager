package duck

import (
	"github.com/fatjan/gogomanager/internal/models"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewDepartmentRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetByID(id int) (*models.Duck, error) {
	duck := &models.Duck{}
	err := r.db.Get(duck, "SELECT id, name FROM public.ducks WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return duck, nil
}

func (r *repository) GetAll() ([]*models.Duck, error) {
	ducks := make([]*models.Duck, 0)
	err := r.db.Get(ducks, "SELECT * from public.ducks")
	if err != nil {
		return nil, err
	}

	return ducks, nil
}
