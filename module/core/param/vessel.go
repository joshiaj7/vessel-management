package param

type CreateVessel struct {
	Name      string
	NACCSCode string
}

type UpdateVessel struct {
	ID        string
	Name      string
	NACCSCode string
}

type ListVessels struct {
	Name    string
	OwnerID string
	Limit   int
	Offset  int
}

type GetVessel struct {
	ID        string
	NACCSCode string
}
