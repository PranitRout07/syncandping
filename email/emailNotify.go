package email

import "github.com/PranitRout07/syncandping/internal/structure"

type EmailNotify struct {
	Email   string
	Passkey string
}

func NewEmailNotify(email string, passkey string) *EmailNotify {
	return &EmailNotify{
		Email:   email,
		Passkey: passkey,
	}
}

func (e *EmailNotify) SendMessage(msg string,label string) (*structure.Resp, error) {
	return nil, nil
}
