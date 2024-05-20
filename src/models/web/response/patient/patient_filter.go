package patientResponse

import "time"

type PatientFilter struct {
	IdentityNumber int64
	PhoneNumber    string
	Name           string
	BirthDate      string
	Gender         string
	CreatedAt      time.Time
}
