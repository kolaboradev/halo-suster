package userResponse

type UserNurseLogin struct {
	UserId      string `json:"userId"`
	Nip         int64  `json:"nip"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}
