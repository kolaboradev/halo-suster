package patientRequest

type PatientFilter struct {
	IdentityNumber *int64
	Limit          int
	Offset         int
	Name           string
	PhoneNumber    string
	CreatedAt      string
}
