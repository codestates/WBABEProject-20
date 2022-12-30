package model

import (
	"strings"

	"github.com/google/uuid"
)

func CreateUUID() string {
	uuidValue := uuid.New()
	uuid := strings.Replace(uuidValue.String(), "-", "", -1)
	return uuid
}
