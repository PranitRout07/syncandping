package main

import (
	"os"

	"github.com/PranitRout07/syncandping/internal/structure"
	"github.com/PranitRout07/syncandping/webhook"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	os.Getenv("slack_webhook")
	w := webhook.NewWebhook("https://hooks.slack.com/services/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")
	csm := structure.CustomMessage{
		Message: "Hii",
		Label: "Normal",
	}

	resp,err := w.SendMessage(csm)
	if err!=nil{
		//
	}
	_ = resp
}
