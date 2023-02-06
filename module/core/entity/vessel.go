package entity

import "time"

type Vessel struct {
	ID        int       `db:"id" json:"id"`
	OwnerID   int       `db:"owner_id" json:"owner_id"`
	Name      string    `db:"name" json:"name"`
	NACCSCode string    `db:"naccs_code" json:"naccs_code"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
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
