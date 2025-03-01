package webhook

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/PranitRout07/syncandping/internal/logger"
	"github.com/PranitRout07/syncandping/internal/structure"
)

type WebhookNotify struct {
	webhook string
}

func NewWebhook(webhook string) *WebhookNotify {
	return &WebhookNotify{
		webhook: webhook,
	}
}

func (w *WebhookNotify) SendMessage(customMessage structure.CustomMessage) (*structure.Resp, error) {

	body, _ := json.Marshal(map[string]string{
		"text": customMessage.Message,
	})
	responseBody := bytes.NewBuffer(body)

	req, err := http.NewRequest("POST", w.webhook, responseBody)
	if err != nil {
		logger.CustomLogger(logger.Logs{Resp: "", Message: customMessage.Message, PingType: "webhook", Label: customMessage.Label, Err: err})
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.CustomLogger(logger.Logs{Resp: "", Message: customMessage.Message, PingType: "webhook", Label: customMessage.Label, Err: err})
		return nil, err
	}
	defer resp.Body.Close()
	
	logger.CustomLogger(logger.Logs{Resp: "Success", Message: customMessage.Message, PingType: "webhook", Label: customMessage.Label, Err: nil})
	return &structure.Resp{Message: "Success", Status: resp.StatusCode}, nil
}
