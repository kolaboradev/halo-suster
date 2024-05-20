package recordEntity

import "time"

type Record struct {
	Id             string
	IdentityNumber int64
	Symptoms       string
	Medications    string
	CreatedBy      string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
