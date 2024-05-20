package userRequest

type (
	UserNurseEdit struct {
		Nip    string `validate:"required,min=13,max=15,nip_nurse"`
		Name   string `validate:"required,min=5,max=50"`
		UserId string
	}

	UserNurseEditParse struct {
		Nip  *int   `json:"nip"`
		Name string `json:"name"`
	}
)
