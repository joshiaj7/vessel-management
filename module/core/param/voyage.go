package param

type CreateVoyage struct {
	VesselID        string
	Source          string
	Destination     string
	CurrentLocation string
}

type UpdateVoyage struct {
	ID              string
	CurrentLocation string
	State           string
}

type ListVoyages struct {
	Limit  int
	Offset int
}

type GetVoyage struct {
	ID string
}
