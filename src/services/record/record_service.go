package recordService

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/kolaboradev/halo-suster/src/exceptions"
	"github.com/kolaboradev/halo-suster/src/helper"
	recordEntity "github.com/kolaboradev/halo-suster/src/models/entities/record"
	recordRequest "github.com/kolaboradev/halo-suster/src/models/web/request/record"
	recordResponse "github.com/kolaboradev/halo-suster/src/models/web/response/record"
	patientRepository "github.com/kolaboradev/halo-suster/src/repositories/patient"
	recordRepository "github.com/kolaboradev/halo-suster/src/repositories/record"
	userRepository "github.com/kolaboradev/halo-suster/src/repositories/user"
)

type RecordService struct {
	DB          *sql.DB
	validator   *validator.Validate
	recordRepo  recordRepository.RecordRepositoryInterface
	patientRepo patientRepository.PatientRepositoryInterface
	userRepo    userRepository.UserRepositoryInterface
}

func NewRecordService(db *sql.DB, validator *validator.Validate, rr recordRepository.RecordRepositoryInterface, pr patientRepository.PatientRepositoryInterface, ur userRepository.UserRepositoryInterface) RecordServiceInterface {
	return &RecordService{
		DB:          db,
		validator:   validator,
		recordRepo:  rr,
		patientRepo: pr,
		userRepo:    ur,
	}
}

func (service *RecordService) Create(ctx context.Context, request recordRequest.Record) {
	err := service.validator.Struct(request)
	helper.ErrorIfPanic(err)

	tx, err := service.DB.Begin()
	helper.ErrorIfPanic(err)
	defer helper.RollbackOrCommit(tx)

	if !service.patientRepo.GetByIdentityNumberBool(ctx, tx, request.IdentityNumber) {
		panic(exceptions.NewNotFoundError("identityNumber is not exist"))
	}

	id := helper.UUIDStr()
	timeNow := helper.TimeISO8601()

	record := recordEntity.Record{
		Id:             id,
		IdentityNumber: request.IdentityNumber,
		Symptoms:       request.Symptoms,
		Medications:    request.Medications,
		CreatedBy:      request.CreatedBy,
		CreatedAt:      timeNow,
		UpdatedAt:      timeNow,
	}
	service.recordRepo.Save(ctx, tx, record)
}

func (service *RecordService) GetAll(ctx context.Context, filters recordRequest.RecordFilter) []recordResponse.RecordFilter {
	tx, err := service.DB.Begin()
	helper.ErrorIfPanic(err)
	defer helper.RollbackOrCommit(tx)

	records := service.recordRepo.GetAll(ctx, tx, filters)

	var recordResponses []recordResponse.RecordFilter
	if len(records) == 0 {
		return recordResponses
	}

	for _, value := range records {
		user, _ := service.userRepo.GetUserById(ctx, tx, value.CreatedBy)
		userInt, err := strconv.Atoi(user.Nip)
		helper.ErrorIfPanic(err)

		userResponse := recordResponse.User{
			Nip:    userInt,
			Name:   user.Name,
			UserId: user.Id,
		}
		patient, _ := service.patientRepo.GetByIdentityNumber(ctx, tx, value.IdentityNumber)

		patienResponse := recordResponse.Patient{
			IdentityNumber:      patient.IdentityNumber,
			PhoneNumber:         patient.PhoneNumber,
			Name:                patient.Name,
			BirthDate:           patient.BirthDate,
			Gender:              patient.BirthDate,
			IdentityCardScanImg: patient.IdentityCardScanImage,
		}

		recordResponses = append(recordResponses, recordResponse.RecordFilter{
			IdentityDetail: patienResponse,
			Symptoms:       value.Symptoms,
			Medications:    value.Medications,
			CreatedAt:      value.CreatedAt,
			CreatedBy:      userResponse,
		})

	}
	return recordResponses
}
