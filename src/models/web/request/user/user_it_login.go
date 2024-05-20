package userRequest

type UserItLogin struct {
	Nip      string `json:"nip" validate:"required,numeric,min=13,max=15,nip_it"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}

type UserItLoginParse struct {
	Nip      *int   `json:"nip"`
	Password string `json:"password"`
}
