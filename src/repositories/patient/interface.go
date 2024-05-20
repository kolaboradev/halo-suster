package patientRepository

import (
	"context"
	"database/sql"

	patientEntity "github.com/kolaboradev/halo-suster/src/models/entities/patient"
	patientRequest "github.com/kolaboradev/halo-suster/src/models/web/request/patient"
)

type PatientRepositoryInterface interface {
	Save(ctx context.Context, tx *sql.Tx, patient *patientEntity.Patient) *patientEntity.Patient
	GetByIdentityNumberBool(ctx context.Context, tx *sql.Tx, identityNumber int64) bool
	GetAll(ctx context.Context, tx *sql.Tx, filters patientRequest.PatientFilter) []patientEntity.Patient
	GetByIdentityNumber(ctx context.Context, tx *sql.Tx, id int64) (patientEntity.Patient, error)
}
