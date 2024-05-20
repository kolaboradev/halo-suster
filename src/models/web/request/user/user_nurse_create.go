package userRequest

type (
	UserNurseCreateParse struct {
		Nip                   *int   `json:"nip"`
		Name                  string `json:"name"`
		IdentityCardScanImage string `json:"identityCardScanImg"`
	}

	UserNurseCreate struct {
		Nip                   string `validate:"required,min=13,max=15,nip_nurse"`
		Name                  string `validate:"required,min=5,max=50"`
		IdentityCardScanImage string `validate:"required,url_image"`
	}
)
