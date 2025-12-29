package adapter

import (
	"fmt"
	"proxy-example/domain"
)

// AppServer is the real subject handling requests.
type AppServer struct {
	logger domain.Logger
}

// NewAppServer builds an application server.
func NewAppServer(logger domain.Logger) *AppServer {
	return &AppServer{logger: logger}
}

func (a *AppServer) HandleRequest(url, method string) (int, string) {
	a.logger.Log(fmt.Sprintf("[AppServer] Handling %s %s", method, url))
	return 200, "OK"
}
