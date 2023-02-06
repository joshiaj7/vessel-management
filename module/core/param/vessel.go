package param

type CreateVessel struct {
	Name      string `json:"name"`
	OwnerID   int    `json:"owner_id"`
	NACCSCode string `json:"naccs_code"`
}

type UpdateVessel struct {
	ID        int    `json:"id"`
	OwnerID   int    `json:"owner_id"`
	Name      string `json:"name"`
	NACCSCode string `json:"naccs_code"`
}

type ListVessels struct {
	Name    string
	OwnerID int
	Limit   int
	Offset  int
}

type GetVessel struct {
	ID int `required:"true"`
}
