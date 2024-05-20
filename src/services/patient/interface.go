package patientService

import (
	patientRequest "github.com/kolaboradev/halo-suster/src/models/web/request/patient"
	patientResponse "github.com/kolaboradev/halo-suster/src/models/web/response/patient"
	"golang.org/x/net/context"
)

type PatientServiceInterface interface {
	Create(ctx context.Context, request patientRequest.PatientCreate) patientResponse.PatientCreate
	GetAll(ctx context.Context, filters patientRequest.PatientFilter) []patientResponse.PatientFilter
}
