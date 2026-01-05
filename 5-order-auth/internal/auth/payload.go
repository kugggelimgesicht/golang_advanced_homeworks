package auth

type LoginResponse struct {
	Token string `json:"token"`
}
type RegisterResponse struct {
	Token string `json:"token"`
}
type VerificationRequest struct {
	SessionId string `json:"sessionId"`
	Code      int    `json:"code"`
}
type SendOTPRequest struct {
	Phone string `json:"phone" validate:"required,min=10,max=15"`
}
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}
