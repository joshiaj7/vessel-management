package param

type CreateShipment struct {
	VesselID        string
	Source          string
	Destination     string
	CurrentLocation string
}

type UpdateShipment struct {
	CurrentLocation string
	State           string
}

type ListShipments struct {
	IDs    []string
	Limit  string
	Offset string
}

type GetShipment struct {
	ID string
}
