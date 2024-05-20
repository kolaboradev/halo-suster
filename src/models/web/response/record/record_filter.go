package recordResponse

import "time"

type (
	Patient struct {
		IdentityNumber      int64  `json:"identityNumber"`
		PhoneNumber         string `json:"phoneNumber"`
		Name                string `json:"name"`
		BirthDate           string `json:"birthDate"`
		Gender              string `json:"gender"`
		IdentityCardScanImg string `json:"identityCardImg"`
	}
	User struct {
		Nip    int    `json:"nip"`
		Name   string `json:"name"`
		UserId string `json:"userId"`
	}

	RecordFilter struct {
		IdentityDetail Patient   `json:"identityDetail"`
		Symptoms       string    `json:"symptoms"`
		Medications    string    `json:"medications"`
		CreatedAt      time.Time `json:"createdAt"`
		CreatedBy      User      `json:"createdBy"`
	}
)
