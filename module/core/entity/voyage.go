package entity

import "time"

type Voyage struct {
	ID                   string     `db:"id"`
	VesselID             string     `db:"vessel_id"`
	Source               string     `db:"source"`
	Destination          string     `db:"destination"`
	CurrentLocation      string     `db:"current_location"`
	State                string     `db:"state"`
	EstimatedArrivalTime *time.Time `db:"estimated_arrival_time"`
	DockedAt             *time.Time `db:"docked_at"`
	DepartedAt           *time.Time `db:"departed_at"`
	ArrivedAt            *time.Time `db:"arrived_at"`
	CreatedAt            time.Time  `db:"created_at"`
	UpdatedAt            time.Time  `db:"updated_at"`
}

func (v *Voyage) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"ID":                   v.ID,
		"VesselID":             v.VesselID,
		"Source":               v.Source,
		"Destination":          v.Destination,
		"CurrentLocation":      v.CurrentLocation,
		"State":                v.State,
		"EstimatedArrivalTime": v.EstimatedArrivalTime,
		"DockedAt":             v.DockedAt,
		"DepartedAt":           v.DepartedAt,
		"ArrivedAt":            v.ArrivedAt,
		"CreatedAt":            v.CreatedAt,
		"UpdatedAt":            v.UpdatedAt,
	}
}
