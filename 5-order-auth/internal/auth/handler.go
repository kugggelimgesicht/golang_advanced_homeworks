package auth

import (
	"net/http"
	"validation-api/configs"
	"validation-api/pkg/jwt"
	"validation-api/pkg/req"
	"validation-api/pkg/res"
)

type AuthHandler struct {
	*configs.Config
	*AuthService
}
type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/verify", handler.VerifyCode())
	router.HandleFunc("POST /auth/send_code", handler.SendOTP())
}

func (handler *AuthHandler) VerifyCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[VerificationRequest](&w, r)
		if err != nil {
			return
		}
		user, err := handler.UserRepository.FindBySessionId(body.SessionId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if user.Code == body.Code {
			token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{Phone: user.Phone})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			data := LoginResponse{
				Token: token,
			}
			res.Json(w, data, http.StatusOK)
		}

	}
}

func (handler *AuthHandler) SendOTP() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[SendOTPRequest](&w, r)
		if err != nil {
			return
		}
		sessionId, err := handler.AuthService.SendOTP(body.Phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		res.Json(w, sessionId, http.StatusOK)
	}
}
