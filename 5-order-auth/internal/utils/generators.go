package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/google/uuid"
)

func GenerateSMSCode() string {
	maxInt := big.NewInt(1000000)
	n, _ := rand.Int(rand.Reader, maxInt)
	return fmt.Sprintf("%x", n.Int64())
}

func GenerateSessionID() string {
	return uuid.New().String()
}
