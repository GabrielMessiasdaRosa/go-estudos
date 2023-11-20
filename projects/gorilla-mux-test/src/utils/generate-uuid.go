package utils

import (
	"github.com/google/uuid"
)

func GenerateUUID() string {
	uuid := uuid.New()
	return uuid.String()
}

func DeuBoa() string {
	deuBoa := "Deu boa!"
	return deuBoa
}