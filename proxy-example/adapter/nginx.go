package adapter

import (
	"fmt"
	"proxy-example/domain"
)

type AppServer struct{}

func (a *AppServer) HandleRequest(url, method string) (int, string) {
	return 200, "OK"
}

type Nginx struct {
	Application       *AppServer
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func NewNginx() *Nginx {
	return &Nginx{
		Application:       &AppServer{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *Nginx) HandleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Forbidden"
	}
	return n.Application.HandleRequest(url, method)
}

func (n *Nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] >= n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url]++
	return true
}
