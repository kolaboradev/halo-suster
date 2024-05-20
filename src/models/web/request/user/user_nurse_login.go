package userRequest

type (
	UserNurseLogin struct {
		Nip      string `json:"nip" validate:"required,numeric,min=13,max=15,nip_nurse"`
		Password string `json:"password" validate:"required,min=5,max=33"`
	}
	UserNurseParse struct {
		Nip      *int   `json:"nip"`
		Password string `json:"password"`
	}
)
