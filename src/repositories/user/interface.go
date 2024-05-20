package userRepository

import (
	"context"
	"database/sql"

	userEntity "github.com/kolaboradev/halo-suster/src/models/entities/user"
	userRequest "github.com/kolaboradev/halo-suster/src/models/web/request/user"
)

type UserRepositoryInterface interface {
	SaveIt(ctx context.Context, tx *sql.Tx, user *userEntity.User) *userEntity.User
	GetUserItByNip(ctx context.Context, tx *sql.Tx, nip string) (userEntity.User, error)
	SaveNurse(ctx context.Context, tx *sql.Tx, user *userEntity.User) *userEntity.User
	GetUserNurseByNip(ctx context.Context, tx *sql.Tx, nip string) (userEntity.User, error)
	GetAllUser(ctx context.Context, tx *sql.Tx, filters userRequest.UserFilter) []userEntity.User
	EditNurseById(ctx context.Context, tx *sql.Tx, user userEntity.User)
	GetNurseByIdBool(ctx context.Context, tx *sql.Tx, id string) bool
	DeleteNurseById(ctx context.Context, tx *sql.Tx, id string)
	GetNurseById(ctx context.Context, tx *sql.Tx, id string) (userEntity.User, error)
	AddNurseAccess(ctx context.Context, tx *sql.Tx, user userEntity.User)
	GetUserById(ctx context.Context, tx *sql.Tx, id string) (userEntity.User, error)
}
