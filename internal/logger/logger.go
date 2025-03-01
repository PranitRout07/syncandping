package logger

import (
	"fmt"


	"golang.org/x/exp/slog"
)

type Logs struct {
	Resp     string
	Message  string
	PingType string
	Label    string
	Err      error
}

func NewLogger()*Logs{
	return &Logs{}
}

func CustomLogger(logs Logs) {
	
	slog.Info(fmt.Sprintf("{ Resp: %s , Message: %s , PingType: %s , Label: %s, Error: %v }",logs.Resp,logs.Message,logs.PingType,logs.Label,logs.Err))
}