package userRequest

type UserItCreate struct {
	Nip      string `json:"nip" validate:"required,numeric,min=13,max=15,nip_it"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}

type UserItCreateParse struct {
	Nip      int    `json:"nip"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
