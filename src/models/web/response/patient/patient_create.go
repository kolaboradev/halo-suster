package patientResponse

type PatientCreate struct {
	IdentityNumber int64  `json:"identityNumber"`
	Name           string `json:"name"`
}
