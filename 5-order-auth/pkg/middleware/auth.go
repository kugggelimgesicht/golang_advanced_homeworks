package middleware

import (
	"context"
	"net/http"
	"strings"
	"validation-api/configs"
	"validation-api/pkg/jwt"
)

type key string

const (
	ContextPhoneKey key = "PhoneKey"
)

func writeUnauthed(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}
func IsAuthorized(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			writeUnauthed(w)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		_, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		ctx := context.WithValue(r.Context(), ContextPhoneKey, data.Phone)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
