package adapter

import (
	"fmt"
	"proxy-example/domain"
)

type AppServer struct {
	logger domain.Logger
}

func NewAppServer(logger domain.Logger) *AppServer {
	return &AppServer{logger: logger}
}

func (a *AppServer) HandleRequest(url, method string) (int, string) {
	a.logger.Log(fmt.Sprintf("[AppServer] Handling %s %s", method, url))
	return 200, "OK"
}
