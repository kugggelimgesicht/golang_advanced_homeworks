package utils

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func GenerateSMSCode() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9000) + 1000
}

func GenerateSessionID() string {
	return uuid.New().String()
}
