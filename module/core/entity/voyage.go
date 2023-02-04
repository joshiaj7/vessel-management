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