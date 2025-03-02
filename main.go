package main

import (
	"os"

	"github.com/PranitRout07/syncandping/email"
	"github.com/PranitRout07/syncandping/pkg/structure"
	// "github.com/PranitRout07/syncandping/webhook"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// wh := os.Getenv("slack_webhook")
	// w := webhook.NewWebhook(wh)
	// csm := structure.CustomMessage{
	// 	Message: "Hii",
	// 	Label: "Normal",
	// }

	// resp,err := w.SendMessage(csm)
	// if err!=nil{
	// 	//
	// }
	// _ = resp

	fromEmail := os.Getenv("from_email")
	toEmail := os.Getenv("to_email")
	passkey := os.Getenv("passkey")

	e := email.NewEmailNotify(fromEmail,toEmail,passkey)
	csm := structure.CustomMessage{
		Message: "Testing email notify",
		Label: "Normal",
	}
	resp,err := e.SendMessage(csm)
	if err!=nil{
		//
	}
	_ = resp
}
