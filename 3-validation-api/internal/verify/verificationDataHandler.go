package verify

import (
	"encoding/json"
	"os"
)

func LoadVerification() (*VerificationRecord, error) {
	data, err := os.ReadFile("data/verification.json")
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var record VerificationRecord
	if err := json.Unmarshal(data, &record); err != nil {
		return nil, err
	}

	return &record, nil
}

func SaveVerification(record *VerificationRecord) error {
	if err := os.MkdirAll("data", 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(record, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("data/verification.json", data, 0644)
}
func deleteVerification() error {
	return os.Remove("data/verification.json")
}
