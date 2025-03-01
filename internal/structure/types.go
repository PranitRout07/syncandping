package structure

import "github.com/PranitRout07/syncandping/internal/logger"

type Resp struct {
	Message string
	Status  int
}

type Logger struct {
	*logger.Logs
}

type CustomMessage struct{
	Message string 	`json:"msg"`
	Label string 	`json:"label,optional"`
}
