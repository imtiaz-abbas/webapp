package utilities

import (
	"github.com/satori/go.uuid"
)

func GenerateID() uuid.UUID {
	return uuid.Must(uuid.NewV4())
}
