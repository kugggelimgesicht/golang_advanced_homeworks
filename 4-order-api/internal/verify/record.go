package verify

type VerificationRecord struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}
