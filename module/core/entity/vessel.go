package entity

import "time"

type Vessel struct {
	ID        string    `db:"id"`
	OwnerID   string    `db:"owner_id"`
	Name      string    `db:"name"`
	NACCSCode string    `db:"naccs_code"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (v *Vessel) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"ID":        v.ID,
		"OwnerID":   v.OwnerID,
		"Name":      v.Name,
		"NACCSCode": v.NACCSCode,
		"CreatedAt": v.CreatedAt,
		"UpdatedAt": v.UpdatedAt,
	}
}
