package auth

import (
	"validation-api/internal/user"
	"validation-api/internal/utils"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) SendOTP(phone string) (string, error) {
	existingUser, _ := service.UserRepository.FindByPhone(phone)
	if existingUser != nil {
		user, err := service.UserRepository.Update(&user.User{
			Phone:     phone,
			SessionId: utils.GenerateSessionID(),
			Code:      utils.GenerateSMSCode(),
		})
		if err != nil {
			return "", err
		}
		return user.SessionId, nil
	}
	user := &user.User{
		Phone:     phone,
		SessionId: utils.GenerateSessionID(),
		Code:      utils.GenerateSMSCode(),
	}
	_, err := service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}
	return user.SessionId, nil
}
