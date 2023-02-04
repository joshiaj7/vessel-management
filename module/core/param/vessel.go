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
	IDs    []string
	Name   string
	Limit  string
	Offset string
}

type GetVessel struct {
	ID        string
	NACCSCode string
}
