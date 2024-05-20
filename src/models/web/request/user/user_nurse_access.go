package userRequest

type (
	UserNurseAccess struct {
		UserId   string
		Password string `json:"password" validate:"required,min=5,max=33"`
	}
	UserNurseAccessParse struct {
		Password string `json:"password"`
	}
)
