package patientEntity

import "time"

type Patient struct {
	Id                    string
	IdentityNumber        int64
	PhoneNumber           string
	Name                  string
	BirthDate             string
	Gender                string
	IdentityCardScanImage string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
