package userResponse

import "time"

type UserFilter struct {
	UserId    string    `json:"userId"`
	Nip       int       `json:"nip"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}
