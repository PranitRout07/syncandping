package email

import (
	"fmt"
	"net/smtp"

	"github.com/PranitRout07/syncandping/internal/logger"
	"github.com/PranitRout07/syncandping/internal/structure"
)

type EmailNotify struct {
	FromEmail string
	ToEmail   string
	Passkey   string
}

const (
	SMTP_HOST = "smtp.gmail.com"
	SMTP_PORT = "587"
)

func NewEmailNotify(fromEmail, toEmail, passkey string) *EmailNotify {
	return &EmailNotify{
		FromEmail: fromEmail,
		ToEmail:   toEmail,
		Passkey:   passkey,
	}
}

func (e *EmailNotify) SendMessage(customMessage structure.CustomMessage) (*structure.Resp, error) {

	to := []string{
		e.ToEmail,
	}

	auth := smtp.PlainAuth("", e.FromEmail, e.Passkey, SMTP_HOST)

	subject := "Subject: Message\n"
	contentType := "MIME-Version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	var msg string
	if customMessage.Label != "" {
		msg = fmt.Sprintf("{ Label: %s }", customMessage.Label) + " " + customMessage.Message
	} else {
		msg = customMessage.Message
	}
	message := []byte(subject + contentType + msg)

	err := smtp.SendMail(SMTP_HOST+":"+SMTP_PORT, auth, e.FromEmail, to, message)
	if err != nil {
		logger.CustomLogger(logger.Logs{Resp: "Failed", Message: customMessage.Message, PingType: "email", Label: customMessage.Label, Err: err})
		return nil, err
	}
	logger.CustomLogger(logger.Logs{Resp: "Success", Message: customMessage.Message, PingType: "email", Label: customMessage.Label, Err: nil})
	return nil, nil
}
