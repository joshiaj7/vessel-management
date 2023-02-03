package entity

import "time"

type Shipment struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	NaccsCode string    `db:"naccs_code"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
