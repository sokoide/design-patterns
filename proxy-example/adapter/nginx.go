package adapter

import (
	"fmt"
	"proxy-example/domain"
)

type Nginx struct {
	application       domain.Server
	maxAllowedRequest int
	rateLimiter       map[string]int
	logger            domain.Logger
}

func NewNginx(app domain.Server, logger domain.Logger) *Nginx {
	return &Nginx{
		application:       app,
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
		logger:            logger,
	}
}

func (n *Nginx) HandleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		n.logger.Log(fmt.Sprintf("[Nginx] Blocked request to %s (Rate Limit Exceeded)", url))
		return 403, "Forbidden"
	}
	n.logger.Log(fmt.Sprintf("[Nginx] Forwarding request to %s", url))
	return n.application.HandleRequest(url, method)
}

func (n *Nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] >= n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url]++
	return true
}
