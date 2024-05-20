package userService

import (
	"context"

	userRequest "github.com/kolaboradev/halo-suster/src/models/web/request/user"
	userResponse "github.com/kolaboradev/halo-suster/src/models/web/response/user"
)

type UserServiceInterface interface {
	CreateIt(ctx context.Context, request userRequest.UserItCreate) userResponse.UserItCreate
	LoginIt(ctx context.Context, request userRequest.UserItLogin) userResponse.UserItLogin
	CreateNurse(ctx context.Context, request userRequest.UserNurseCreate) userResponse.UserNurseCreate
	LoginNurse(ctx context.Context, request userRequest.UserNurseLogin) userResponse.UserNurseLogin
	GetAllUsers(ctx context.Context, filters userRequest.UserFilter) []userResponse.UserFilter
	EditNurseById(ctx context.Context, request userRequest.UserNurseEdit)
	DeleteNurseById(ctx context.Context, id string)
	AddNurseAccess(ctx context.Context, request userRequest.UserNurseAccess)
}
