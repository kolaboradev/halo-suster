package userEntity

import (
	"database/sql"
	"time"
)

type User struct {
	Id                    string
	Nip                   string
	Name                  string
	Password              sql.NullString
	Role                  string
	IdentityCardScanImage sql.NullString
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
