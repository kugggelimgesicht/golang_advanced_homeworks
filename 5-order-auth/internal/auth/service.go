package auth

import (
	"log"
	"validation-api/internal/user"
	"validation-api/internal/utils"

	"gorm.io/gorm"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func sendSMSCode(phone string) (int, error) {
	code := utils.GenerateSMSCode()

	// имитация отправки SMS
	log.Printf("Send SMS to %s with code: %s\n", phone, code)

	return code, nil
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) SendOTP(phone string) (string, error) {
	code, err := sendSMSCode(phone)
	if err != nil {
		return "", err
	}
	existingUser, _ := service.UserRepository.FindByPhone(phone)
	if existingUser != nil {
		user, err := service.UserRepository.Update(&user.User{
			Model:     gorm.Model{ID: existingUser.ID},
			Phone:     phone,
			SessionId: utils.GenerateSessionID(),
			Code:      code,
		})
		if err != nil {
			return "", err
		}
		return user.SessionId, nil
	}
	user := &user.User{
		Phone:     phone,
		SessionId: utils.GenerateSessionID(),
		Code:      code,
	}
	_, err = service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}
	return user.SessionId, nil
}
