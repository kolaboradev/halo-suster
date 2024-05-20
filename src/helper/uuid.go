package helper

import "github.com/google/uuid"

func UUIDStr() string {
	id := uuid.New()
	return id.String()
}
