package req

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomHash(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
