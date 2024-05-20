package patientRequest

type PatientCreate struct {
	IdentityNumber      int64  `validate:"required,number,identity_number"`
	PhoneNumber         string `validate:"required,min=10,max=15,startswith=+62"`
	Name                string `validate:"required,min=3,max=30"`
	BirthDate           string `validate:"required,datetime=2006-01-02"`
	Gender              string `validate:"required,gender"`
	IdentityCardScanImg string `validate:"required,url_image"`
}

// ,datetime=2006-01-02
