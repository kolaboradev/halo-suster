package userRequest

type UserFilter struct {
	UserId    string
	Limit     int
	Offset    int
	Name      string
	Nip       *int
	Role      string
	CreatedAt string
}
