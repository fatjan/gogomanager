package models

import "time"

type Department struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	ManagerID int       `db:"manager_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
