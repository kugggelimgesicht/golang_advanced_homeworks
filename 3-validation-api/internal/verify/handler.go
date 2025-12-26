package verify

import (
	"errors"
	"fmt"
	"net/http"
	"net/smtp"
	"validation-api/configs"
	"validation-api/pkg/req"
	"validation-api/pkg/res"

	"github.com/jordan-wright/email"
)

type VerifHandler struct {
	*configs.Config
}
type VerifHandlerDeps struct {
	*configs.Config
}

func NewVerifHandler(router *http.ServeMux, deps VerifHandlerDeps) {
	handler := &VerifHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VerifHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[SendEmailRequest](w, r)
		if err != nil {
			return
		}
		conf := handler.Config
		hash, _ := req.GenerateRandomHash(16)
		e := email.NewEmail()
		e.From = "Jordan Wright <test@gmail.com>"
		e.To = []string{body.Email}
		e.Bcc = []string{"test_bcc@example.com"}
		e.Cc = []string{"test_cc@example.com"}
		e.Subject = "Awesome Subject"
		e.HTML = []byte(fmt.Sprintf(
			`<a>verify your email via this link: http://localhost:8081/verify/%s</a>`,
			hash))
		err = e.Send(conf.Address.Address, smtp.PlainAuth("", conf.Email.Email, conf.Password.Password, "smtp.gmail.com"))
		if err != nil {
			res.Json(w, err, http.StatusBadRequest)
			return
		}

		res.Json(w, "verification email sent", http.StatusOK)
		record := &VerificationRecord{
			Email: body.Email,
			Hash:  hash,
		}

		if err = SaveVerification(record); err != nil {
			res.Json(w, err, http.StatusInternalServerError)
			return
		}
	}
}

func (handler *VerifHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ok, err := handler.verifyLogic(r)
		if err != nil {
			res.Json(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !ok {
			res.Json(w, "verification failed", http.StatusNotFound)
			return
		}

		res.Json(w, "verification email sent", http.StatusOK)
	}
}

func (handler *VerifHandler) verifyLogic(r *http.Request) (bool, error) {
	hash := r.PathValue("hash")
	if hash == "" {
		return false, errors.New("hash is required")
	}

	record, err := LoadVerification()
	if err != nil {
		return false, err
	}
	if record == nil {
		return false, nil
	}
	if record.Hash != hash {
		_ = deleteVerification()
		return false, errors.New("invalid verification link")
	}
	_ = deleteVerification()
	return true, nil
}
