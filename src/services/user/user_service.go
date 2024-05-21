package userService

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/kolaboradev/halo-suster/src/exceptions"
	"github.com/kolaboradev/halo-suster/src/helper"
	userEntity "github.com/kolaboradev/halo-suster/src/models/entities/user"
	userRequest "github.com/kolaboradev/halo-suster/src/models/web/request/user"
	userResponse "github.com/kolaboradev/halo-suster/src/models/web/response/user"
	userRepository "github.com/kolaboradev/halo-suster/src/repositories/user"
)

type UserService struct {
	DB             *sql.DB
	validator      *validator.Validate
	userRepository userRepository.UserRepositoryInterface
}

func NewUserService(db *sql.DB, validator *validator.Validate, userRepo userRepository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		DB:             db,
		validator:      validator,
		userRepository: userRepo,
	}
}

func (service *UserService) CreateIt(ctx context.Context, request userRequest.UserItCreate) userResponse.UserItCreate {
	err := service.validator.Struct(request)
	helper.ErrorIfPanic(err)

	tx, err := service.DB.Begin()
	helper.ErrorIfPanic(err)
	defer helper.RollbackOrCommit(tx)

	_, err = service.userRepository.GetUserItByNip(ctx, tx, request.Nip)
	if err == nil {
		panic(exceptions.NewConflictError("nip already use"))
	}

	id := helper.UUIDStr()
	timeNow := helper.TimeISO8601()

	passwordHash := helper.HashPassword(request.Password)

	user := userEntity.User{
		Id:        id,
		Nip:       request.Nip,
		Name:      request.Name,
		Password:  sql.NullString{String: passwordHash},
		Role:      "it",
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	service.userRepository.SaveIt(ctx, tx, &user)

	token := helper.GenerateTokenUserJWT(user)
	nipInt, err := strconv.Atoi(user.Nip)
	helper.ErrorIfPanic(err)

	return userResponse.UserItCreate{
		UserId:      user.Id,
		Nip:         int64(nipInt),
		Name:        user.Name,
		AccessToken: token,
	}
}

func (service *UserService) LoginIt(ctx context.Context, request userRequest.UserItLogin) userResponse.UserItLogin {
	if !helper.IsItNiP(request.Nip) {
		panic(exceptions.NewNotFoundError("user is not from IT not with 615"))
	}

	err := service.validator.Struct(request)
	helper.ErrorIfPanic(err)

	tx, err := service.DB.Begin()
	helper.ErrorIfPanic(err)
	defer helper.RollbackOrCommit(tx)

	user, err := service.userRepository.GetUserItByNip(ctx, tx, request.Nip)
	if err != nil {
		panic(exceptions.NewNotFoundError("user it not found"))
	}

	passwordValid := helper.CompareHashPassword(user.Password.String, request.Password)
	if !passwordValid {
		panic(exceptions.NewBadRequestError("password is wrong"))
	}

	token := helper.GenerateTokenUserJWT(user)

	nipInt, err := strconv.Atoi(user.Nip)
	helper.ErrorIfPanic(err)
	return userResponse.UserItLogin{
		UserId:      user.Id,
		Nip:         int64(nipInt),
		Name:        user.Name,
		AccessToken: token,
	}
}

func (service *UserService) CreateNurse(ctx context.Context, request userRequest.UserNurseCreate) userResponse.UserNurseCreate {
	err := service.validator.Struct(request)
	helper.ErrorIfPanic(err)

	tx, err := service.DB.Begin()
	helper.ErrorIfPanic(err)
	defer helper.RollbackOrCommit(tx)

	_, err = service.userRepository.GetUserNurseByNip(ctx, tx, request.Nip)
	if err == nil {
		panic(exceptions.NewConflictError("nip nurse must exist"))
	}

	id := helper.UUIDStr()
	timeNow := helper.TimeISO8601()
	user := userEntity.User{
		Id:                    id,
		Nip:                   request.Nip,
		Name:                  request.Name,
		Role:                  "nurse",
		IdentityCardScanImage: sql.NullString{String: request.IdentityCardScanImage},
		CreatedAt:             timeNow,
		UpdatedAt:             timeNow,
	}
	service.userRepository.SaveNurse(ctx, tx, &user)

	return userResponse.UserNurseCreate{
		UserId: user.Id,
		Nip:    user.Nip,
		Name:   user.Name,
	}
}

func (service *UserService) LoginNurse(ctx context.Context, request userRequest.UserNurseLogin) userResponse.UserNurseLogin {
	if !helper.IsNurseNiP(request.Nip) {
		panic(exceptions.NewNotFoundError("user is not from IT not with 303"))
	}

	err := service.validator.Struct(request)
	helper.ErrorIfPanic(err)

	tx, err := service.DB.Begin()
	helper.ErrorIfPanic(err)
	defer helper.RollbackOrCommit(tx)

	user, err := service.userRepository.GetUserNurseByNip(ctx, tx, request.Nip)
	if !user.Password.Valid {
		panic(exceptions.NewBadRequestError("user is not having access"))
	}
	if err != nil {
		panic(exceptions.NewNotFoundError("user nurse not found"))
	}

	token := helper.GenerateTokenUserJWT(user)

	nipNum, err := strconv.Atoi(user.Nip)
	helper.ErrorIfPanic(err)

	return userResponse.UserNurseLogin{
		UserId:      user.Id,
		Nip:         int64(nipNum),
		Name:        user.Name,
		AccessToken: token,
	}

}

func (service *UserService) GetAllUsers(ctx context.Context, filters userRequest.UserFilter) []userResponse.UserFilter {
	tx, err := service.DB.Begin()
	helper.ErrorIfPanic(err)
	defer helper.RollbackOrCommit(tx)

	users := service.userRepository.GetAllUser(ctx, tx, filters)

	var userResponses []userResponse.UserFilter
	for _, value := range users {
		nipInt, err := strconv.Atoi(value.Nip)
		helper.ErrorIfPanic(err)
		user := userResponse.UserFilter{
			UserId:    value.Id,
			Nip:       nipInt,
			Name:      value.Name,
			CreatedAt: value.CreatedAt,
		}
		userResponses = append(userResponses, user)
	}
	return userResponses
}

func (service *UserService) EditNurseById(ctx context.Context, request userRequest.UserNurseEdit) {
	if !helper.IsNurseNiP(request.Nip) {
		panic(exceptions.NewBadRequestError("nip not starts with 303"))
	}
	err := service.validator.Struct(request)
	helper.ErrorIfPanic(err)

	tx, err := service.DB.Begin()
	helper.ErrorIfPanic(err)
	defer helper.RollbackOrCommit(tx)

	nurseMustExist, err := service.userRepository.GetUserById(ctx, tx, request.UserId)
	if err != nil {
		panic(exceptions.NewNotFoundError("nurse not found"))
	}
	if err == nil {
		if !helper.IsNurseNiP(nurseMustExist.Nip) {
			panic(exceptions.NewNotFoundError("nurse not found"))
		}
	}

	// if !nurseMustExist {
	// 	panic(exceptions.NewNotFoundError("nurse not found"))
	// }

	nurse, err := service.userRepository.GetUserItByNip(ctx, tx, request.Nip)
	if err == nil {
		if nurse.Id != request.UserId {
			panic(exceptions.NewConflictError("nip exists"))
		}
	}

	timeNow := helper.TimeISO8601()

	nurseEdit := userEntity.User{
		Id:        request.UserId,
		Nip:       request.Nip,
		Name:      request.Name,
		UpdatedAt: timeNow,
	}

	service.userRepository.EditNurseById(ctx, tx, nurseEdit)

}

func (service *UserService) DeleteNurseById(ctx context.Context, id string) {
	tx, err := service.DB.Begin()
	helper.ErrorIfPanic(err)
	defer helper.RollbackOrCommit(tx)

	nurse, err := service.userRepository.GetNurseById(ctx, tx, id)

	if err != nil {
		panic(exceptions.NewNotFoundError("nurse not found"))
	}
	if !helper.IsNurseNiP(nurse.Nip) {
		panic(exceptions.NewNotFoundError("nip not starts with 303"))
	}

	service.userRepository.DeleteNurseById(ctx, tx, id)
}

func (service *UserService) AddNurseAccess(ctx context.Context, request userRequest.UserNurseAccess) {
	err := service.validator.Struct(request)
	helper.ErrorIfPanic(err)

	tx, err := service.DB.Begin()
	helper.ErrorIfPanic(err)
	defer helper.RollbackOrCommit(tx)

	nurse, err := service.userRepository.GetNurseById(ctx, tx, request.UserId)

	if err != nil {
		panic(exceptions.NewNotFoundError("nurse not found"))
	}
	if !helper.IsNurseNiP(nurse.Nip) {
		panic(exceptions.NewNotFoundError("nip not starts with 303"))
	}

	user := userEntity.User{
		Id:       request.UserId,
		Password: sql.NullString{String: request.Password},
	}

	service.userRepository.AddNurseAccess(ctx, tx, user)
}
