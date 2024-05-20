package userResponse

type UserNurseCreate struct {
	UserId string `json:"userId"`
	Nip    string `json:"nip"`
	Name   string `json:"name"`
}
