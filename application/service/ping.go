package service

import (
	"bingo-example/infrastructure/util"
)

type PingService struct {
}

// Execute 执行指令
func (s *PingService) Execute(f string) (*util.Response, error) {
	var t string
	var data interface{}

	switch f {
	case "ping":
		t, data = s.Pong()
	}

	return util.NewResponse(t, data), nil
}

// Pong 回复心跳ping
func (s *PingService) Pong() (string, interface{}) {
	return "ping", "pong"
}
