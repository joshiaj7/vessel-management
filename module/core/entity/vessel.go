package entity

import "time"

type Vessel struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	NACCSCode string    `db:"naccs_code"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
