package recordRequest

type Record struct {
	IdentityNumber int64  `validate:"required,number,identity_number"`
	Symptoms       string `validate:"required,min=1,max=2000"`
	Medications    string `validate:"required,min=1,max=2000"`
	CreatedBy      string
}

type RecordParse struct {
	IdentityNumber int64  `validate:"required,number,identity_number"`
	Symptoms       string `validate:"required,min=1,max=2000"`
	Medications    string `validate:"required,min=1,max=2000"`
}
