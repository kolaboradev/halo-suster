package patientResponse

import "time"

type PatientFilter struct {
	IdentityNumber int64     `json:"identityNumber"`
	PhoneNumber    string    `json:"phoneNumber"`
	Name           string    `json:"name"`
	BirthDate      string    `json:"birthDate"`
	Gender         string    `json:"gender"`
	CreatedAt      time.Time `json:"createdAt"`
}
