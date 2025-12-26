package verify

import (
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
		record := &VerificationRecord{
			Email: body.Email,
			Hash:  hash,
		}

		if err = SaveVerification(record); err != nil {
			res.Json(w, err, http.StatusInternalServerError)
			return
		}
		if err != nil {
			res.Json(w, err, http.StatusInternalServerError)
			return
		}

		res.Json(w, "verification email sent", http.StatusOK)
	}
}

func (handler *VerifHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		if hash == "" {
			res.Json(w, "hash is required", http.StatusBadRequest)
			return
		}

		record, err := LoadVerification()
		if err != nil {
			res.Json(w, err, http.StatusInternalServerError)
			return
		}
		if record.Hash != hash {
			_ = deleteVerification()
			res.Json(w, "invalid verification link", http.StatusNotFound)
			return
		}
	}
}
