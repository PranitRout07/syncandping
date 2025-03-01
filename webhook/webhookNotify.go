package webhook

import "github.com/PranitRout07/syncandping/structure"

type WebhookNotify struct {
	webhook string
}

func NewWebhook(webhook string) *WebhookNotify {
	return &WebhookNotify{
		webhook: webhook,
	}
}

func (w *WebhookNotify) SendMessage(msg string) (*structure.Resp, error) {

	return nil, nil
}
