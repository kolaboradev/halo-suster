package patientService

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/kolaboradev/halo-suster/src/exceptions"
	"github.com/kolaboradev/halo-suster/src/helper"
	patientEntity "github.com/kolaboradev/halo-suster/src/models/entities/patient"
	patientRequest "github.com/kolaboradev/halo-suster/src/models/web/request/patient"
	patientResponse "github.com/kolaboradev/halo-suster/src/models/web/response/patient"
	patientRepository "github.com/kolaboradev/halo-suster/src/repositories/patient"
	"golang.org/x/net/context"
)

type PatientService struct {
	DB          *sql.DB
	validator   *validator.Validate
	patientRepo patientRepository.PatientRepositoryInterface
}

func NewPatientService(db *sql.DB, validator *validator.Validate, pr patientRepository.PatientRepositoryInterface) PatientServiceInterface {
	return &PatientService{
		DB:          db,
		validator:   validator,
		patientRepo: pr,
	}
}

func (service *PatientService) Create(ctx context.Context, request patientRequest.PatientCreate) patientResponse.PatientCreate {
	err := service.validator.Struct(request)
	helper.ErrorIfPanic(err)

	tx, err := service.DB.Begin()
	helper.ErrorIfPanic(err)
	defer helper.RollbackOrCommit(tx)

	if service.patientRepo.GetByIdentityNumberBool(ctx, tx, request.IdentityNumber) {
		panic(exceptions.NewConflictError("identity number already exist"))
	}

	id := helper.UUIDStr()
	timeNow := helper.TimeISO8601()

	patient := patientEntity.Patient{
		Id:                    id,
		IdentityNumber:        request.IdentityNumber,
		PhoneNumber:           request.PhoneNumber,
		Name:                  request.Name,
		BirthDate:             request.BirthDate,
		Gender:                request.Gender,
		IdentityCardScanImage: request.IdentityCardScanImg,
		CreatedAt:             timeNow,
		UpdatedAt:             timeNow,
	}

	service.patientRepo.Save(ctx, tx, &patient)

	return patientResponse.PatientCreate{
		IdentityNumber: patient.IdentityNumber,
		Name:           patient.Name,
	}
}

func (service *PatientService) GetAll(ctx context.Context, filters patientRequest.PatientFilter) []patientResponse.PatientFilter {
	tx, err := service.DB.Begin()
	helper.ErrorIfPanic(err)
	defer helper.RollbackOrCommit(tx)

	users := service.patientRepo.GetAll(ctx, tx, filters)

	var usersResponses []patientResponse.PatientFilter
	for _, value := range users {
		user := patientResponse.PatientFilter{
			IdentityNumber: value.IdentityNumber,
			PhoneNumber:    value.PhoneNumber,
			Name:           value.Name,
			BirthDate:      value.BirthDate,
			Gender:         value.Gender,
			CreatedAt:      value.CreatedAt,
		}
		usersResponses = append(usersResponses, user)
	}
	return usersResponses
}
